// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"testing"
	"time"

	"github.com/Altoros/gosigma/data"
	"github.com/Altoros/gosigma/mock"
)

var mockEndpoint string

func init() {
	mock.Start()
	mockEndpoint = mock.Endpoint("")
}

func newDataServer() *data.Server {
	return &data.Server{
		ServerRecord: data.ServerRecord{
			Resource: data.Resource{URI: "uri", UUID: "uuid"},
			Name:     "name",
			Status:   "status",
		},
		Meta: map[string]string{"key1": "value1", "key2": "value2"},
	}
}

func createTestClient(t *testing.T) (*Client, error) {
	cli, err := NewClient(mockEndpoint, mock.TestUser, mock.TestPassword, nil)
	if err != nil {
		return nil, err
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	return cli, nil
}

func TestClientCreate(t *testing.T) {
	check := func(ep, u, p string) {
		cli, err := NewClient(ep, u, p, nil)
		if err == nil || cli != nil {
			t.Errorf("NewClient(%q,%q,%q) must fail", ep, u, p)
		}
		t.Log("OK:", err)
	}

	// endpoint
	check("", "1234", "1234")
	check("1234", "1234", "1234")
	check("https://1234:1234@endpoint.com", "1234", "1234")
	check("https://endpoint.com?xxx", "1234", "1234")
	check("://endpoint.com?xxx", "1234", "1234")

	// auth
	check(mockEndpoint, "", "")
	check(mockEndpoint, "", "1234")
	check(mockEndpoint, "1234", "")

	// OK
	cli, err := createTestClient(t)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
	}
}

type testLog struct{ written int }

func (l *testLog) Log(args ...interface{})                 { l.written += 1 }
func (l *testLog) Logf(format string, args ...interface{}) { l.written += 1 }

func TestClientLogger(t *testing.T) {
	cli, err := NewClient("https://1.0.0.0:2000/api/2.0/", mock.TestUser, mock.TestPassword, nil)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	var log testLog
	cli.Logger(&log)

	cli.ConnectTimeout(100 * time.Millisecond)
	cli.ReadWriteTimeout(100 * time.Millisecond)

	ssf, err := cli.Servers(false)
	if err == nil || ssf != nil {
		t.Error("Servers(false) returned valid result for unavailable endpoint")
		return
	}

	if log.written == 0 {
		t.Error("no writes to log")
	}
}

func TestClientEmptyUUID(t *testing.T) {
	cli, err := createTestClient(t)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	if _, err := cli.Server(""); err != errEmptyUUID {
		t.Error("Server('') must fail with errEmptyUUID")
	}
	if err := cli.StartServer("", nil); err != errEmptyUUID {
		t.Error("StartServer('') must fail with errEmptyUUID")
	}
	if err := cli.StopServer(""); err != errEmptyUUID {
		t.Error("StopServer('') must fail with errEmptyUUID")
	}
}

func TestClientEndpointUnavailableSoft(t *testing.T) {
	cli, err := NewClient(mockEndpoint+"1", mock.TestUser, mock.TestPassword, nil)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	ssf, err := cli.Servers(false)
	if err == nil || ssf != nil {
		t.Error("AllServers(false) returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(false)", err)

	sst, err := cli.Servers(true)
	if err == nil || sst != nil {
		t.Error("AllServers(true) returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(true)", err)

	s, err := cli.Server("uuid")
	if err == nil {
		t.Error("Server() returned valid result with for unavailable endpoint: %#v", s)
		return
	}
	t.Log("OK, Server():", err)

	if _, err := cli.CreateServer(Components{}); err == nil {
		t.Error("CreateServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, CreateServer():", err)

	err = cli.StartServer("uuid", nil)
	if err == nil {
		t.Error("StartServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, StartServer():", err)

	err = cli.StopServer("uuid")
	if err == nil {
		t.Error("StopServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, StopServer():", err)
}

func TestClientEndpointUnavailableHard(t *testing.T) {
	cli, err := NewClient("https://1.0.0.0:2000/api/2.0/", mock.TestUser, mock.TestPassword, nil)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	cli.ConnectTimeout(100 * time.Millisecond)
	cli.ReadWriteTimeout(100 * time.Millisecond)

	ssf, err := cli.Servers(false)
	if err == nil || ssf != nil {
		t.Error("Servers(false) returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK: Servers(false)", err)

	sst, err := cli.Servers(true)
	if err == nil || sst != nil {
		t.Error("Servers(true) returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK: Servers(true)", err)

	s, err := cli.Server("uuid")
	if err == nil {
		t.Error("Server() returned valid result for unavailable endpoint: %#v", s)
		return
	}
	t.Log("OK, Server():", err)

	if _, err := cli.CreateServer(Components{}); err == nil {
		t.Error("CreateServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, CreateServer():", err)

	err = cli.StartServer("uuid", nil)
	if err == nil {
		t.Error("StartServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, StartServer():", err)

	err = cli.StopServer("uuid")
	if err == nil {
		t.Error("StopServer() returned valid result for unavailable endpoint")
		return
	}
	t.Log("OK, StopServer():", err)
}
