// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"io"
)

const (
	// ServerStopped defines constant for stopped instance state
	ServerStopped = "stopped"
	// ServerStarting defines constant for starting instance state
	ServerStarting = "starting"
	// ServerRunning defines constant for running instance state
	ServerRunning = "running"
	// ServerStopping defines constant for stopping instance state
	ServerStopping = "stopping"
	// ServerUnavailable defines constant for unavailable instance state
	ServerUnavailable = "unavailable"
)

// NIC describes properties of network interface card
type NIC struct {
	IPv4 struct {
		Conf string `json:"conf"`
		IP   struct {
			URI  string `json:"resource_uri"`
			UUID string `json:"uuid"`
		} `json:"ip"`
	} `json:"ip_v4_conf"`
}

// Server contains properties of cloud server instance
type Server struct {
	Name   string            `json:"name"`
	URI    string            `json:"resource_uri"`
	Status string            `json:"status"`
	UUID   string            `json:"uuid"`
	Meta   map[string]string `json:"meta"`
	NICs   []NIC             `json:"nics"`
}

// Servers holds collection of cloud server instances
type Servers struct {
	Meta struct {
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Objects []Server `json:"objects"`
}

// ReadServers reads and unmarshalls information about cloud server instances from JSON stream
func ReadServers(r io.Reader) ([]Server, error) {
	var servers Servers
	dec := json.NewDecoder(r)
	for {
		err := dec.Decode(&servers)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return servers.Objects, nil
}

// ReadServer reads and unmarshalls information about single cloud server instance from JSON stream
func ReadServer(r io.Reader) (*Server, error) {
	var server Server
	dec := json.NewDecoder(r)
	for {
		err := dec.Decode(&server)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return &server, nil
}
