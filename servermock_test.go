// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"crypto/tls"
	"flag"
	"net/http"
	"testing"
)

var mock *Servermock = CreateServerMock("2.0")
var tr *http.Transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client *http.Client = &http.Client{Transport: tr}

func get(t *testing.T, url string) {
	_, err := client.Get(url)
	if err != nil {
		t.Errorf("(%T)%v", err, err)
	}
}

type testWriter struct {
	t *testing.T
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.t.Log(string(p))
	return len(p), nil
}

func init() {
	flag.Var(&mock.Log, "log", "")
}

func setup(t *testing.T) {
	if testing.Verbose() {
		mock.Writer = &testWriter{t}
	}
}

func TestServerMockSections(t *testing.T) {
	setup(t)

	check := func(s string) {
		get(t, mock.Endpoint+s)
		if mock.LastSection != s {
			t.Errorf("Section == %s, wants %s", mock.LastSection, s)
		}
	}

	check("capabilities")
	check("drives")
}
