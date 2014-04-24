// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/Altoros/gosigma/data"
)

// A Factory designed to create objects in CloudSigma account
type Factory struct {
	*Client
}

// A ServerConfiguration structure describes configuration parameters for new server.
type ServerConfiguration struct {
	// Name for new server
	Name string
	// Cpu frequency in MHz
	CPU int64
	// Mem capacity in bytes
	Mem int64
	// TemplateDrive UUID to clone drive for new server from.
	TemplateDrive string
	// DriveName for newly cloned drive
	DriveName string
	// VLan UUID to attach newly created server to.
	VLan string
	// VNCPassword for new server
	VNCPassword string
	// SSHPublicKey defines SSH public key for new server
	SSHPublicKey string
	// Description of server
	Description string
}

// CreateServerByParameters creates server from passed parameters
func (f Factory) CreateServerFromConfiguration(p ServerConfiguration) (Server, error) {

	// clone drive
	d, err := f.CloneDrive(p.TemplateDrive, p.DriveName)
	if err != nil {
		return Server{}, err
	}

	// prepare parameters
	var m = make(map[string]interface{})

	// common parameters
	m["name"] = p.Name
	m["cpu"] = p.CPU
	m["mem"] = p.Mem
	m["vnc_password"] = p.VNCPassword

	// meta-information
	var meta = make(map[string]string)
	meta["description"] = p.Description
	meta["ssh_public_key"] = p.SSHPublicKey
	m["meta"] = meta

	// template drive
	var drive = make(map[string]interface{})
	drive["boot_order"] = 1
	drive["dev_channel"] = "0:0"
	drive["device"] = "virtio"
	drive["drive"] = data.MakeResource("drives", d.UUID())
	m["drives"] = []interface{}{drive}

	// nics
	var dhcp = make(map[string]interface{})
	dhcp["ip_v4_conf"] = map[string]string{"conf": "dhcp"}
	dhcp["model"] = "virtio"

	var vlan = make(map[string]interface{})
	vlan["vlan"] = data.MakeResource("vlans", p.VLan)
	vlan["model"] = "virtio"

	m["nics"] = []interface{}{dhcp, vlan}

	// serialize
	bb, err := json.Marshal(m)
	if err != nil {
		return Server{}, nil
	}
	rr := bytes.NewReader(bb)

	// run request
	u := f.endpoint + "servers/"
	r, err := f.https.Post(u, nil, rr)
	if err != nil {
		return Server{}, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(201); err != nil {
		return Server{}, NewError(r, err)
	}

	ss, err := data.ReadServers(r.Body)
	if err != nil {
		return Server{}, err
	}
	if len(ss) == 0 {
		return Server{}, errors.New("no servers in response from endpoint")
	}

	s := Server{f.Client, &ss[0]}

	return s, nil
}

// CreateServerFromJSON creates new server instance(s) from passed JSON
func (f Factory) CreateServerFromJSON(json string) ([]Server, error) {
	u := f.endpoint + "servers/"

	content := strings.NewReader(json)
	r, err := f.https.Post(u, nil, content)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(201); err != nil {
		return nil, NewError(r, err)
	}

	objs, err := data.ReadServers(r.Body)
	if err != nil {
		return nil, err
	}

	servers := make([]Server, len(objs))
	for i := 0; i < len(objs); i++ {
		servers[i] = Server{
			client: f.Client,
			obj:    &objs[i],
		}
	}

	return servers, nil
}

// CloneDrive clones given drive
func (f Factory) CloneDrive(uuid, name string) (Drive, error) {
	// clone the drive
	var params = CloneParams{Name: name}
	dd, err := f.clone(uuid, &params, nil)
	if err != nil {
		return Drive{}, err
	}
	if len(dd) == 0 {
		return Drive{}, errors.New("no drives in response from endpoint")
	}

	// wait for job is done
	d := Drive{f.Client, &dd[0]}
	jj := d.Jobs()
	if len(jj) == 0 {
		return d, nil
	}

	j := jj[0]

	var chStop chan int
	var stop = false

	readWriteTimeout := f.https.GetReadWriteTimeout()
	if readWriteTimeout > 0 {
		chStop = make(chan int)
		go func() {
			select {
			case <-time.After(readWriteTimeout):
				stop = true
			case <-chStop:
				return
			}
		}()
	}

	for !stop && j.Progress() < 100 {
		err := j.Refresh()
		if err != nil {
			return Drive{}, err
		}
	}

	if chStop != nil {
		close(chStop)
	}

	if stop {
		return Drive{}, errors.New("timeout waiting for finish of drive cloning process")
	}

	d.Refresh()

	return d, nil
}
