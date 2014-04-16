// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Altoros/gosigma"
)

//
// Implementation of CloudSigma server mock object for testing purposes.
//
//	Username: test@example.com
//	Password: test
//

const serverBase = "/api/2.0/"

const (
	TestUser     = "test@example.com"
	TestPassword = "test"
)

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
	mux.HandleFunc(makeHandler("servers", serversHandler))

	pServer = httptest.NewUnstartedServer(mux)
	pServer.StartTLS()
}

type handlerType func(http.ResponseWriter, *http.Request)

func makeHandler(name string, f handlerType) (string, handlerType) {
	url := serverBase + name + "/"
	handler := func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()

		if isValidAuth(r) {
			f(rec, r)
		} else {
			rec.WriteHeader(401)
			rec.Write([]byte("401 Unauthorized\n"))
		}

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

func isValidAuth(r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) < 2 {
		return false
	}
	switch s[0] {
	case "Basic":
		return isValidBasicAuth(s[1])
	case "Digest":
		return isValidDigestAuth(s[1])
	}

	return false
}

func isValidBasicAuth(auth string) bool {
	b, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return false
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}
	if pair[0] != TestUser {
		return false
	}
	if pair[1] != TestPassword {
		return false
	}
	return true
}

func isValidDigestAuth(auth string) bool {
	return false
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

	url := Endpoint() + s

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(TestUser, TestPassword)

	client := gosigma.NewHttpsClient(nil)

	return client.Do(req)
}
