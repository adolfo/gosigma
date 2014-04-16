// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"io"
)

const (
	ServerStopped     = "stopped"
	ServerStarting    = "starting"
	ServerRunning     = "running"
	ServerStopping    = "stopping"
	ServerUnavailable = "unavailable"
)

type Nic struct {
	IPv4 struct {
		Conf string `json:"conf"`
		IP   struct {
			URI  string `json:"resource_uri"`
			UUID string `json:"uuid"`
		} `json:"ip"`
	} `json:"ip_v4_conf"`
}

type Server struct {
	Name   string            `json:"name"`
	URI    string            `json:"resource_uri"`
	Status string            `json:"status"`
	UUID   string            `json:"uuid"`
	Meta   map[string]string `json:"meta"`
	NICs   []Nic             `json:"nics"`
}

type Servers struct {
	Meta struct {
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Objects []Server `json:"objects"`
}

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
