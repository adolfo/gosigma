// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/Altoros/gosigma/data"
)

const (
	Kilobyte = 1024
	Megabyte = 1024 * 1024
	Gigabyte = 1024 * 1024 * 1024
	Terabyte = 1024 * 1024 * 1024 * 1024
)

const (
	ModelVirtio = "virtio"
	ModelE1000  = "e1000"
)

// A Components contains information to create new server
type Components struct {
	m map[string]interface{}
}

// SetName sets name for new server. To unset name, call this function with empty string in the name parameter.
func (c *Components) SetName(name string) {
	c.setString("name", name)
}

// SetCPU sets CPU frequency for new server. To unset CPU frequency, call this function with zero in the frequency parameter.
func (c *Components) SetCPU(frequency int64) {
	c.setInt("cpu", frequency)
}

// SetMem sets memory size for new server. To unset this value, call function with zero in the bytes parameter.
func (c *Components) SetMem(bytes int64) {
	c.setInt("mem", bytes)
}

// SetVNCPassword sets VNC password for new server. To unset, call this function with empty string.
func (c *Components) SetVNCPassword(password string) {
	c.setString("vnc_password", password)
}

// SetDescription sets description for new server. To unset, call this function with empty string.
func (c *Components) SetDescription(description string) {
	c.setMeta("description", description)
}

// SetPublicSSHKey sets public SSH key for new server. To unset, call this function with empty string.
func (c *Components) SetSSHPublicKey(description string) {
	c.setMeta("ssh_public_key", description)
}

// AttachDriveData attaches drive to components from drive data.
func (c *Components) AttachDrive(drive Drive, bootOrder int, channel, device string) {
	c.init()

	var dm = make(map[string]interface{})
	if bootOrder > 0 {
		dm["boot_order"] = bootOrder
	}
	if channel != "" {
		dm["dev_channel"] = channel
	}
	if device != "" {
		dm["device"] = device
	}
	if drive.obj != nil && drive.obj.UUID != "" {
		dm["drive"] = drive.obj.Resource
	}

	if len(dm) > 0 {
		dd, _ := c.m["drives"].([]interface{})
		c.m["drives"] = append(dd, dm)
	}
}

// NetworkDHCP4 attaches NIC, configured with IPv4 DHCP
func (c *Components) NetworkDHCP4(model string) {
	c.network4(model, "dhcp", "")
}

// NetworkDHCP4 attaches NIC, configured with IPv4 static address
func (c *Components) NetworkStatic4(model, address string) {
	c.network4(model, "static", address)
}

// NetworkManual4 attaches NIC, configured with IPv4 manual settings
func (c *Components) NetworkManual4(model string) {
	c.network4(model, "manual", "")
}

// NetworkManual4 attaches NIC, configured with private VLan
func (c *Components) NetworkVLan(model, uuid string) {
	c.init()

	var nm = make(map[string]interface{})

	if model != "" {
		nm["model"] = model
	}

	if uuid != "" {
		nm["vlan"] = data.MakeVLanResource(uuid)
	}

	if len(nm) > 0 {
		nics, _ := c.m["nics"].([]interface{})
		c.m["nics"] = append(nics, nm)
	}
}

func (c *Components) network4(model, conf, address string) {
	c.init()

	var nm = make(map[string]interface{})

	if model != "" {
		nm["model"] = model
	}

	var conf4 = map[string]interface{}{"conf": conf}
	if address != "" {
		conf4["ip"] = data.MakeIPResource(address)
	}
	nm["ip_v4_conf"] = conf4

	nics, _ := c.m["nics"].([]interface{})
	c.m["nics"] = append(nics, nm)
}

func (c *Components) init() {
	if c.m == nil {
		c.m = make(map[string]interface{})
	}
}

func (c *Components) setString(name, value string) {
	c.init()
	value = strings.TrimSpace(value)
	if value == "" {
		delete(c.m, name)
	} else {
		c.m[name] = value
	}
}

func (c *Components) setInt(name string, value int64) {
	c.init()
	if value == 0 {
		delete(c.m, name)
	} else {
		c.m[name] = value
	}
}

func (c *Components) setMeta(name, value string) {
	c.init()

	mi := c.m["meta"]
	m, _ := mi.(map[string]string)

	value = strings.TrimSpace(value)
	if value == "" {
		delete(m, name)
		if len(m) == 0 {
			delete(c.m, "meta")
		}
	} else {
		if m == nil {
			m = make(map[string]string)
			c.m["meta"] = m
		}
		m[name] = value
	}
}

func (c Components) marshal() (io.Reader, error) {
	bb, err := json.Marshal(c.m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bb), nil
}

func (c Components) marshalString() (string, error) {
	r, err := c.marshal()
	if err != nil {
		return "", err
	}
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}
