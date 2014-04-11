// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

//
// Private implementation of CloudSigma server mock object for testing purposes.
//

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type Logtype int

func (l Logtype) String() string {
	switch l {
	case LogSection:
		return "section"
	case LogURL:
		return "url"
	case LogDetail:
		return "detail"
	default:
		return "none"
	}
}

func (l *Logtype) Set(v string) error {
	switch v {
	case "section":
		*l = LogSection
	case "url":
		*l = LogURL
	case "detail":
		*l = LogDetail
	default:
		*l = LogNone
	}
	return nil
}

const (
	LogNone    Logtype = 0
	LogSection         = 1
	LogURL             = 2
	LogDetail          = 3
)

type Servermock struct {
	URL         string // https://host:port
	Endpoint    string // https://host:port/api/{version}/
	Log         Logtype
	TLS         *tls.Config
	LastRequest *http.Request
	LastSection string
	Writer      io.Writer
	server      *httptest.Server
}

// Create new servermock object
func CreateServerMock(apiversion string) *Servermock {
	mock := &Servermock{}

	mux := http.NewServeMux()

	base := "/api/" + apiversion + "/"

	mux.HandleFunc(base+"capabilities/", func(w http.ResponseWriter, r *http.Request) {
		mock.handleCapabilities(w, r)
	})
	mux.HandleFunc(base+"drives/", func(w http.ResponseWriter, r *http.Request) {
		mock.handleDrives(w, r)
	})

	server := httptest.NewUnstartedServer(mux)
	server.StartTLS()

	mock.TLS = server.TLS
	mock.URL = server.URL
	mock.Endpoint = mock.URL + base
	mock.server = server

	return mock
}

func (m Servermock) Close() {
	m.server.CloseClientConnections()
	m.server.Close()
}

func (m *Servermock) log(s string, r *http.Request) {
	m.LastSection = s
	m.LastRequest = r

	if m.Writer == nil {
		return
	}

	switch m.Log {
	case LogSection:
		fmt.Fprint(m.Writer, s)
	case LogURL:
		fmt.Fprint(m.Writer, r.RequestURI)
	case LogDetail:
		r.Write(m.Writer)
	}
}

// /capabilities/ section handler
func (m *Servermock) handleCapabilities(w http.ResponseWriter, r *http.Request) {
	m.log("capabilities", r)
}

// /drives/ section handler
func (m *Servermock) handleDrives(w http.ResponseWriter, r *http.Request) {
	m.log("drives", r)
}
