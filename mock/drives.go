// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"

	"github.com/Altoros/gosigma/data"
)

var syncDrives sync.Mutex
var drives = make(map[string]*data.Drive)

func initDrive(d *data.Drive) (*data.Drive, error) {
	if d.UUID == "" {
		uuid, err := GenerateUUID()
		if err != nil {
			return nil, err
		}
		d.UUID = uuid
	}
	if d.Status == "" {
		d.Status = "unmounted"
	}

	return d, nil
}

// AddDrive adds drive instance record under the mock
func AddDrive(d *data.Drive) error {
	d, err := initDrive(d)
	if err != nil {
		return err
	}

	syncDrives.Lock()
	defer syncDrives.Unlock()

	drives[d.UUID] = d

	return nil
}

// AddDrives adds drive instance records under the mock
func AddDrives(dd []data.Drive) []string {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	var result []string
	for _, d := range dd {
		d, err := initDrive(&d)
		if err != nil {
			drives[d.UUID] = d
			result = append(result, d.UUID)
		}
	}
	return result
}

// RemoveDrive removes drive instance record from the mock
func RemoveDrive(uuid string) bool {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	_, ok := drives[uuid]
	delete(drives, uuid)
	return ok
}

// ResetDrives removes all drive instance records from the mock
func ResetDrives() {
	syncDrives.Lock()
	defer syncDrives.Unlock()
	drives = make(map[string]*data.Drive)
}

// SetDriveStatus changes status of server instance in the mock
func SetDriveStatus(uuid, status string) {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	d, ok := drives[uuid]
	if ok {
		d.Status = status
	}
}

// CloneDrive clones specified drive
var ErrNotFound = errors.New("not found")

func CloneDrive(uuid string) (string, error) {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	d, ok := drives[uuid]
	if !ok {
		return "", ErrNotFound
	}

	newUUID, err := GenerateUUID()
	if err != nil {
		return "", err
	}

	var newDrive data.Drive = *d
	newDrive.Resource = *data.MakeDriveResource(newUUID)
	drives[newUUID] = &newDrive

	return newUUID, nil
}

func drivesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		drivesHandlerGet(w, r)
	case "POST":
		drivesHandlerPost(w, r)
	case "DELETE":
		drivesHandlerDelete(w, r)
	}
}

func drivesHandlerGet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	switch path {
	case "/api/2.0/drives":
		handleDrives(w, r)
	case "/api/2.0/drives/detail":
		handleDrivesDetail(w, r, 200, nil)
	default:
		uuid := strings.TrimPrefix(path, "/api/2.0/drives/")
		handleDrive(w, r, 200, uuid)
	}
}

func drivesHandlerPost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/action/")
	uuid := strings.TrimPrefix(path, "/api/2.0/drives/")
	handleDriveAction(w, r, uuid)
}

func drivesHandlerDelete(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	uuid := strings.TrimPrefix(path, "/api/2.0/drives/")
	if ok := RemoveDrive(uuid); !ok {
		h := w.Header()
		h.Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	}
	w.WriteHeader(204)
}

func handleDrives(w http.ResponseWriter, r *http.Request) {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	var dd data.Drives
	dd.Meta.TotalCount = len(drives)
	dd.Objects = make([]data.Drive, 0, len(drives))
	for _, d := range drives {
		var drv data.Drive
		drv.Resource = d.Resource
		drv.Owner = d.Owner
		drv.Status = d.Status
		dd.Objects = append(dd.Objects, drv)
	}

	data, err := json.Marshal(&dd)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func handleDrivesDetail(w http.ResponseWriter, r *http.Request, okcode int, filter []string) {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	var dd data.Drives

	if len(filter) == 0 {
		dd.Meta.TotalCount = len(drives)
		dd.Objects = make([]data.Drive, 0, len(drives))
		for _, d := range drives {
			dd.Objects = append(dd.Objects, *d)
		}
	} else {
		dd.Meta.TotalCount = len(filter)
		dd.Objects = make([]data.Drive, 0, len(filter))
		for _, uuid := range filter {
			if d, ok := drives[uuid]; ok {
				dd.Objects = append(dd.Objects, *d)
			}
		}
	}

	data, err := json.Marshal(&dd)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(okcode)
	w.Write(data)
}

func handleDrive(w http.ResponseWriter, r *http.Request, okcode int, uuid string) {
	syncDrives.Lock()
	defer syncDrives.Unlock()

	h := w.Header()

	d, ok := drives[uuid]
	if !ok {
		h.Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	}

	data, err := json.Marshal(&d)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}

	h.Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(okcode)
	w.Write(data)
}

func handleDriveAction(w http.ResponseWriter, r *http.Request, uuid string) {
	vv := r.URL.Query()

	v, ok := vv["do"]
	if !ok || len(v) < 1 {
		w.WriteHeader(400)
		return
	}

	action := v[0]
	switch action {
	case "clone":
		handleDriveClone(w, r, uuid)
	default:
		w.WriteHeader(400)
	}
}

func handleDriveClone(w http.ResponseWriter, r *http.Request, uuid string) {
	newUUID, err := CloneDrive(uuid)
	if err == ErrNotFound {
		h := w.Header()
		h.Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		w.Write([]byte(jsonNotFound))
		return
	} else if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 " + err.Error()))
		return
	}
	handleDrivesDetail(w, r, 202, []string{newUUID})
}
