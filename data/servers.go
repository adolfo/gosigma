// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import "io"

// IPv4 describes properties of IPv4 address
type IPv4 struct {
	Conf string   `json:"conf"`
	IP   Resource `json:"ip"`
}

// NIC describes properties of network interface card
type NIC struct {
	IPv4  IPv4     `json:"ip_v4_conf"`
	Model string   `json:"model"`
	MAC   string   `json:"mac"`
	VLAN  Resource `json:"vlan"`
}

// ServerDrive describe properties of disk drive
type ServerDrive struct {
	BootOrder int      `json:"boot_order"`
	Channel   string   `json:"dev_channel"`
	Device    string   `json:"device"`
	Drive     Resource `json:"drive"`
}

// ServerRecord contains main properties of cloud server instance
type ServerRecord struct {
	Resource
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ServerRecords holds collection of ServerRecord objects
type ServerRecords struct {
	Meta    Meta           `json:"meta"`
	Objects []ServerRecord `json:"objects"`
}

// Server contains detail properties of cloud server instance
type Server struct {
	ServerRecord
	CPU    int64             `json:"cpu"`
	Mem    int64             `json:"mem"`
	Meta   map[string]string `json:"meta"`
	NICs   []NIC             `json:"nics"`
	Drives []ServerDrive     `json:"drives"`
}

// Servers holds collection of Server objects
type Servers struct {
	Meta    Meta     `json:"meta"`
	Objects []Server `json:"objects"`
}

// ReadServers reads and unmarshalls description of cloud server instances from JSON stream
func ReadServers(r io.Reader) ([]Server, error) {
	var servers Servers
	if err := ReadJson(r, &servers); err != nil {
		return nil, err
	}
	return servers.Objects, nil
}

// ReadServer reads and unmarshalls description of single cloud server instance from JSON stream
func ReadServer(r io.Reader) (*Server, error) {
	var server Server
	if err := ReadJson(r, &server); err != nil {
		return nil, err
	}
	return &server, nil
}
