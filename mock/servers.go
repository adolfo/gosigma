// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/Altoros/gosigma/data"
)

var syncServers sync.Mutex
var servers = make(map[string]*data.Server)

// AddServer adds server instance record under the mock
func AddServer(s *data.Server) {
	syncServers.Lock()
	defer syncServers.Unlock()

	servers[s.UUID] = s
}

// RemoveServer removes server instance record from the mock
func RemoveServer(uuid string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	delete(servers, uuid)
}

// ResetServers removes all server instance records from the mock
func ResetServers() {
	syncServers.Lock()
	defer syncServers.Unlock()
	servers = make(map[string]*data.Server)
}

// SetServerStatus changes status of server instance in the mock
func SetServerStatus(uuid, status string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	s, ok := servers[uuid]
	if ok {
		s.Status = status
	}
}

// URLs:
// /api/2.0/servers
// /api/2.0/servers/detail/
// /api/2.0/servers/{uuid}/
func serversHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/api/2.0/servers" {
		handleServers(w, r)
		return
	}
	if path == "/api/2.0/servers/detail" {
		handleServersDetail(w, r)
		return
	}

	uuid := strings.TrimPrefix(path, "/api/2.0/servers/")
	handleServer(w, r, uuid)
}

func handleServers(w http.ResponseWriter, r *http.Request) {
	syncServers.Lock()
	defer syncServers.Unlock()

	var ss data.ServerRecords
	ss.Meta.TotalCount = len(servers)
	ss.Objects = make([]data.ServerRecord, 0, len(servers))
	for _, s := range servers {
		ss.Objects = append(ss.Objects, s.ServerRecord)
	}

	data, err := json.Marshal(&ss)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h := w.Header()
	h["Content-Type"] = append(h["Content-Type"], "application/json")
	h["Content-Type"] = append(h["Content-Type"], "charset=utf-8")
	w.Write(data)
}

func handleServersDetail(w http.ResponseWriter, r *http.Request) {
	syncServers.Lock()
	defer syncServers.Unlock()

	var ss data.Servers
	ss.Meta.TotalCount = len(servers)
	ss.Objects = make([]data.Server, 0, len(servers))
	for _, s := range servers {
		ss.Objects = append(ss.Objects, *s)
	}

	data, err := json.Marshal(&ss)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h := w.Header()
	h["Content-Type"] = append(h["Content-Type"], "application/json")
	h["Content-Type"] = append(h["Content-Type"], "charset=utf-8")
	w.Write(data)
}

func handleServer(w http.ResponseWriter, r *http.Request, uuid string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	s, ok := servers[uuid]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("404 Not found\n"))
		return
	}

	data, err := json.Marshal(&s)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h := w.Header()
	h["Content-Type"] = append(h["Content-Type"], "application/json")
	h["Content-Type"] = append(h["Content-Type"], "charset=utf-8")
	w.Write(data)
}
