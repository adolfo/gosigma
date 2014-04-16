// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"io"
)

type Server struct {
	Name   string `json:"name"`
	URI    string `json:"resource_uri"`
	Status string `json:"status"`
	UUID   string `json:"uuid"`
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
