// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

//
// Implementation of CloudSigma server mock object for testing purposes.
//

const serverBase = "/api/2.0/"

var pServer *httptest.Server

// Start mock server for testing CloudSigma endpoint communication.
// If server is already started, this function does nothing.
func Start() {
	if IsStarted() {
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc(makeHandler("capabilities", capsHandler))
	mux.HandleFunc(makeHandler("drives", drivesHandler))

	pServer = httptest.NewUnstartedServer(mux)
	pServer.StartTLS()
}

type handlerType func(http.ResponseWriter, *http.Request)

func makeHandler(name string, f handlerType) (string, handlerType) {
	url := serverBase + name + "/"
	handler := func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		f(rec, r)

		recordJournal(name, r, rec)

		hdr := w.Header()
		for k, v := range rec.HeaderMap {
			hdr[k] = v
		}

		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
	}
	return url, handler
}

// Check the mock server is started
func IsStarted() bool {
	if pServer == nil {
		return false
	}
	return true
}

// Stop mock server.
// Panic if server is not started.
func Stop() {
	pServer.CloseClientConnections()
	pServer.Close()
	pServer = nil
}

// Endpoint of mock server, represented as string in form 'https://host:port/api/{version}/'.
// Panic if server is not started.
func Endpoint() string {
	return pServer.URL + serverBase
}

// Request mock server for given URL
func Request(s string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := Endpoint() + s
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
