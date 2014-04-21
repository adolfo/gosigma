// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

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

const jsonNotFound = `[{
		"error_point": null,
	 	"error_type": "notexist",
	 	"error_message": "notfound"
}]`

const jsonStartFailed = `[{
		"error_point": null,
		"error_type": "permission",
		"error_message": "Cannot start guest in state \"started\". Guest should be in state \"stopped\""
}]`

const jsonStopFailed = `[{
		"error_point": null,
		"error_type": "permission",
		"error_message": "Cannot stop guest in state \"stopped\". Guest should be in state \"['started', 'running_legacy']\""
}]`

const jsonActionSuccess = `{
		"action": "%s",
		"result": "success",
		"uuid": "%s"
}`

// URLs:
// /api/2.0/servers
// /api/2.0/servers/detail/
// /api/2.0/servers/{uuid}/
func serversHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		serversHandlerGet(w, r)
	case "POST":
		serversHandlerPost(w, r)
	}
}

func serversHandlerGet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	switch path {
	case "/api/2.0/servers":
		handleServers(w, r)
	case "/api/2.0/servers/detail":
		handleServersDetail(w, r)
	default:
		uuid := strings.TrimPrefix(path, "/api/2.0/servers/")
		handleServer(w, r, uuid)
	}
}

func serversHandlerPost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/action/")
	uuid := strings.TrimPrefix(path, "/api/2.0/servers/")
	handleServerAction(w, r, uuid)
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
	h.Set("Content-Type", "application/json; charset=utf-8")
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
	h.Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func handleServer(w http.ResponseWriter, r *http.Request, uuid string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	h := w.Header()

	s, ok := servers[uuid]
	if !ok {
		h.Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	}

	data, err := json.Marshal(&s)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h.Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func handleServerAction(w http.ResponseWriter, r *http.Request, uuid string) {
	vv := r.URL.Query()

	v, ok := vv["do"]
	if !ok || len(v) < 1 {
		w.WriteHeader(400)
		return
	}

	action := v[0]
	switch action {
	case "start":
		handleServerStart(w, r, uuid)
	case "stop":
		handleServerStop(w, r, uuid)
	default:
		w.WriteHeader(400)
		return
	}
}

func handleServerStart(w http.ResponseWriter, r *http.Request, uuid string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")

	s, ok := servers[uuid]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	}

	if !strings.HasPrefix(s.Status, "stopped") {
		w.WriteHeader(403)
		w.Write([]byte(jsonStartFailed))
		return
	}

	s.Status = "starting"
	go func() {
		syncServers.Lock()
		defer syncServers.Unlock()
		<-time.After(300 * time.Millisecond)
		s.Status = "running"
	}()

	w.Write([]byte(fmt.Sprintf(string(jsonActionSuccess), "start", s.UUID)))
}

func handleServerStop(w http.ResponseWriter, r *http.Request, uuid string) {
	syncServers.Lock()
	defer syncServers.Unlock()

	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")

	s, ok := servers[uuid]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	}

	if !strings.HasPrefix(s.Status, "running") {
		w.WriteHeader(403)
		w.Write([]byte(jsonStopFailed))
		return
	}

	s.Status = "stopping"
	go func() {
		syncServers.Lock()
		defer syncServers.Unlock()
		<-time.After(300 * time.Millisecond)
		s.Status = "stopped"
	}()

	w.Write([]byte(fmt.Sprintf(jsonActionSuccess, "stop", s.UUID)))
}
