// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"testing"

	"github.com/Altoros/gosigma/data"
)

func TestComponentsName(t *testing.T) {
	var c Components
	c.SetName("test")
	if v, ok := c.m["name"]; !ok || v != "test" {
		t.Errorf("invalid SetName(\"test\"), ok=%v, v=%v", ok, v)
	}
	c.SetName("")
	if v, ok := c.m["name"]; ok || v != nil {
		t.Errorf("invalid SetName(\"\"), ok=%v, v=%v", ok, v)
	}
}

func TestComponentsCPU(t *testing.T) {
	var c Components
	c.SetCPU(2000)
	if v, ok := c.m["cpu"]; !ok || v != int64(2000) {
		t.Errorf("invalid SetCPU(2000), ok=%v, v=%v", ok, v)
	}
	c.SetCPU(0)
	if v, ok := c.m["cpu"]; ok || v != nil {
		t.Errorf("invalid SetCPU(0), ok=%v, v=%v", ok, v)
	}
}

func TestComponentsMem(t *testing.T) {
	var c Components
	c.SetMem(5368709120)
	if v, ok := c.m["mem"]; !ok || v != int64(5368709120) {
		t.Errorf("invalid SetMem(5368709120), ok=%v, v=%v", ok, v)
	}
	c.SetMem(0)
	if v, ok := c.m["mem"]; ok || v != nil {
		t.Errorf("invalid SetMem(0), ok=%v, v=%v", ok, v)
	}
}

func TestComponentsVNCPassword(t *testing.T) {
	var c Components
	c.SetVNCPassword("test")
	if v, ok := c.m["vnc_password"]; !ok || v != "test" {
		t.Errorf("invalid SetVNCPassword(\"test\"), ok=%v, v=%v", ok, v)
	}
	c.SetVNCPassword("")
	if v, ok := c.m["vnc_password"]; ok || v != nil {
		t.Errorf("invalid SetVNCPassword(\"\"), ok=%v, v=%v", ok, v)
	}
}

func TestComponentsDescription(t *testing.T) {
	var c Components

	c.SetDescription("test")

	mi, ok := c.m["meta"]
	m := mi.(map[string]string)
	if !ok || m == nil {
		t.Errorf("invalid SetDescription(\"test\"), ok=%v, v=%v", ok, m)
	} else {
		if v, ok := m["description"]; !ok || v != "test" {
			t.Errorf("invalid SetDescription(\"test\"), ok=%v, v=%v", ok, v)
		}
	}

	c.SetDescription("")
	mi, ok = c.m["meta"]
	if ok || mi != nil {
		t.Errorf("invalid SetDescription(\"\"), ok=%v, v=%v", ok, m)
	}
}

func TestComponentsSSHPublicKey(t *testing.T) {
	var c Components

	c.SetSSHPublicKey("test")

	mi, ok := c.m["meta"]
	m := mi.(map[string]string)
	if !ok || m == nil {
		t.Errorf("invalid SetDescription(\"test\"), ok=%v, v=%v", ok, m)
	} else {
		if v, ok := m["ssh_public_key"]; !ok || v != "test" {
			t.Errorf("invalid SetDescription(\"test\"), ok=%v, v=%v", ok, v)
		}
	}

	c.SetSSHPublicKey("")
	mi, ok = c.m["meta"]
	if ok || mi != nil {
		t.Errorf("invalid SetDescription(\"\"), ok=%v, v=%v", ok, m)
	}
}

func TestComponentsAttachDrive(t *testing.T) {
	var c Components
	var d = data.ServerDrive{
		BootOrder: 1,
		Channel:   "0:0",
		Device:    "virtio",
		Drive:     data.MakeDriveResource("uuid"),
	}

	c.AttachDrive(d)

	di, ok := c.m["drives"]
	if !ok || di == nil {
		t.Errorf("invalid AttachDrive call")
		return
	}

	dd, ok := di.([]interface{})
	if !ok || dd == nil {
		t.Errorf("invalid AttachDrive call")
		return
	}

	if len(dd) != 1 {
		t.Errorf("invalid AttachDrive call")
		return
	}

	mi := dd[0]
	mm, ok := mi.(map[string]interface{})
	if !ok || mm == nil {
		t.Errorf("invalid AttachDrive call")
		return
	}

	if v, ok := mm["boot_order"]; !ok || v != 1 {
		t.Errorf("invalid AttachDrive call: ok=%v, v=%v, wants 1", ok, v)
	}
	if v, ok := mm["dev_channel"]; !ok || v != "0:0" {
		t.Errorf("invalid AttachDrive call: ok=%v, v=%v, wants '0:0'", ok, v)
	}
	if v, ok := mm["device"]; !ok || v != "virtio" {
		t.Errorf("invalid AttachDrive call: ok=%v, v=%v, wants 'virtio'", ok, v)
	}

	dui, ok := mm["drive"]
	if !ok || dui == nil {
		t.Errorf("invalid AttachDrive call: ok=%v, dui=%v", ok, dui)
	}

	du, ok := dui.(data.Resource)
	if !ok {
		t.Errorf("invalid AttachDrive call: ok=%v, du=%v", ok, du)
	}

	duv := data.MakeDriveResource("uuid")
	if du != duv {
		t.Errorf("invalid AttachDrive call: resource=%v, wants %v", du, duv)
	}
}

func TestComponentsAttachEmptyDrive(t *testing.T) {
	var c Components
	var d = data.ServerDrive{}

	c.AttachDrive(d)

	di, ok := c.m["drives"]
	if ok || di != nil {
		t.Errorf("invalid AttachDrive call")
		return
	}
}

func TestComponentsAttachDHCPNic(t *testing.T) {
	var c Components

	var n = data.NIC{}
	n.Model = "virtio"
	n.IPv4.Conf = "dhcp"

	c.AttachNIC(n)

	if s, err := c.marshalString(); err != nil {
		t.Error(err)
	} else if v := `{"nics":[{"ip_v4_conf":{"conf":"dhcp"},"model":"virtio"}]}`; s != v {
		t.Errorf("invalid AttachNIC, returned `%s`, wants `%s`", s, v)
	}
}

func TestComponentsAttachStaticNic(t *testing.T) {
	var c Components

	var n = data.NIC{}
	n.Model = "virtio"
	n.IPv4.Conf = "static"
	n.IPv4.IP = data.MakeIPResource("ipaddr")

	c.AttachNIC(n)

	if s, err := c.marshalString(); err != nil {
		t.Error(err)
	} else if v := `{"nics":[{"ip_v4_conf":{"conf":"static","ip":"ipaddr"},"model":"virtio"}]}`; s != v {
		t.Errorf("invalid AttachNIC, returned `%s`, wants `%s`", s, v)
	}
}

func TestComponentsAttachManualNic(t *testing.T) {
	var c Components

	var n = data.NIC{}
	n.Model = "virtio"
	n.IPv4.Conf = "manual"

	c.AttachNIC(n)

	if s, err := c.marshalString(); err != nil {
		t.Error(err)
	} else if v := `{"nics":[{"ip_v4_conf":{"conf":"manual"},"model":"virtio"}]}`; s != v {
		t.Errorf("invalid AttachNIC, returned `%s`, wants `%s`", s, v)
	}
}

func TestComponentsAttachVLanNic(t *testing.T) {
	var c Components

	var n = data.NIC{}
	n.Model = "virtio"
	n.VLAN = data.MakeVLanResource("vlanuuid")

	c.AttachNIC(n)

	if s, err := c.marshalString(); err != nil {
		t.Error(err)
	} else if v := `{"nics":[{"model":"virtio","vlan":"vlanuuid"}]}`; s != v {
		t.Errorf("invalid AttachNIC, returned `%s`, wants `%s`", s, v)
	}
}

type noMarshal int

func (noMarshal) MarshalJSON() ([]byte, error) {
	return nil, errors.New("error")
}

func TestComponentsMarshalEmpty(t *testing.T) {
	var c Components
	c.init()
	c.m["bad"] = noMarshal(0)
	s, err := c.marshalString()
	if err == nil {
		t.Error(err, s)
	} else {
		t.Log(err)
	}
}
