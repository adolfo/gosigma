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

func createTestClient() (*Client, error) {
	return NewClient(mockEndpoint, mock.TestUser, mock.TestPassword, nil)
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
	cli, err := createTestClient()
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
	}
}

func TestClientEndpointUnavailableSoft(t *testing.T) {
	cli, err := NewClient(mockEndpoint+"1", mock.TestUser, mock.TestPassword, nil)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	ssf, err := cli.AllServers(false)
	if err == nil || ssf != nil {
		t.Error("AllServers(false) returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(false)", err)

	sst, err := cli.AllServers(true)
	if err == nil || sst != nil {
		t.Error("AllServers(true) returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(true)", err)

	s, err := cli.Server("uuid")
	if err == nil || s != nil {
		t.Error("Server() returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK, Server():", err)
}

func TestClientEndpointUnavailableHard(t *testing.T) {
	cli, err := NewClient("https://1.0.0.0:2000/api/2.0/", mock.TestUser, mock.TestPassword, nil)
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	cli.ConnectTimeout(100 * time.Millisecond)
	cli.ReadWriteTimeout(100 * time.Millisecond)

	ssf, err := cli.AllServers(false)
	if err == nil || ssf != nil {
		t.Error("AllServers(false) returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(false)", err)

	sst, err := cli.AllServers(true)
	if err == nil || sst != nil {
		t.Error("AllServers(true) returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK: AllServers(true)", err)

	s, err := cli.Server("uuid")
	if err == nil || s != nil {
		t.Error("Server() returned valid result with unavailable endpoint")
		return
	}
	t.Log("OK, Server():", err)
}

func TestClientAllServersEmpty(t *testing.T) {
	mock.ResetServers()

	cli, err := createTestClient()
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
		return
	}

	check := func(detail bool) {
		servers, err := cli.AllServers(detail)
		if err != nil {
			t.Error(err)
		}
		if len(servers) > 0 {
			t.Errorf("%v", servers)
		}
	}

	check(false)
	check(true)
}

func TestClientAllServers(t *testing.T) {
	mock.ResetServers()

	ds := data.Server{
		ServerRecord: data.ServerRecord{
			Name:   "name",
			URI:    "uri",
			Status: "status",
			UUID:   "uuid",
		},
		Meta: map[string]string{"key1": "value1", "key2": "value2"},
	}
	mock.AddServer(&ds)

	cli, err := createTestClient()
	if err != nil {
		t.Error(err)
		return
	}

	servers, err := cli.AllServers(true)
	if err != nil {
		t.Error(err)
		return
	}

	if len(servers) != 1 {
		t.Errorf("invalid len: %v", servers)
		return
	}

	s := servers[0]

	if s.String() == "" {
		t.Error("Empty string representation")
		return
	}

	checkv := func(v, wants string) {
		if v != wants {
			t.Errorf("value %s, wants %s", v, wants)
		}
	}
	checkv(s.Name(), "name")
	checkv(s.URI(), "uri")
	checkv(s.Status(), "status")
	checkv(s.UUID(), "uuid")

	checkg := func(s Server, k, wants string) {
		if v, ok := s.Get(k); !ok || v != wants {
			t.Errorf("value of Get(%q) = %q, %v; wants %s", k, v, ok, wants)
		}
	}
	checkg(s, "key1", "value1")
	checkg(s, "key2", "value2")

	// refresh
	ds.Name = "name1"
	ds.URI = "uri1"
	ds.Status = "status1"
	ds.Meta["key1"] = "value11"
	ds.Meta["key2"] = "value22"
	ds.Meta["key3"] = "value33"
	if err := s.Refresh(); err != nil {
		t.Error(err)
		return
	}
	checkv(s.Name(), "name1")
	checkv(s.URI(), "uri1")
	checkv(s.Status(), "status1")
	checkg(s, "key1", "value11")
	checkg(s, "key2", "value22")
	checkg(s, "key3", "value33")

	// failed refresh
	mock.ResetServers()
	if err := s.Refresh(); err == nil {
		t.Error("Server refresh must fail")
		return
	}

	mock.ResetServers()
}

func TestClientServer(t *testing.T) {
	mock.ResetServers()

	ds := data.Server{
		ServerRecord: data.ServerRecord{
			Name:   "name",
			URI:    "uri",
			Status: "status",
			UUID:   "uuid",
		},
		Meta: map[string]string{"key1": "value1", "key2": "value2"},
	}
	mock.AddServer(&ds)

	cli, err := createTestClient()
	if err != nil {
		t.Error(err)
		return
	}

	if s, err := cli.Server(""); err == nil || s != nil {
		t.Error(err)
		return
	}

	s, err := cli.Server("uuid")
	if err != nil || s == nil {
		t.Error(err)
		return
	}

	if s.String() == "" {
		t.Error("Empty string representation")
	}

	checkv := func(v, wants string) {
		if v != wants {
			t.Errorf("value %s, wants %s", v, wants)
		}
	}
	checkv(s.Name(), "name")
	checkv(s.URI(), "uri")
	checkv(s.Status(), "status")
	checkv(s.UUID(), "uuid")

	checkg := func(s Server, k, wants string) {
		if v, ok := s.Get(k); !ok || v != wants {
			t.Errorf("value of Get(%q) = %q, %v; wants %s", k, v, ok, wants)
		}
	}
	checkg(*s, "key1", "value1")
	checkg(*s, "key2", "value2")

	// refresh
	ds.Name = "name1"
	ds.URI = "uri1"
	ds.Status = "status1"
	ds.Meta["key1"] = "value11"
	ds.Meta["key2"] = "value22"
	ds.Meta["key3"] = "value33"
	if err := s.Refresh(); err != nil {
		t.Error(err)
	}
	checkv(s.Name(), "name1")
	checkv(s.URI(), "uri1")
	checkv(s.Status(), "status1")
	checkg(*s, "key1", "value11")
	checkg(*s, "key2", "value22")
	checkg(*s, "key3", "value33")

	// failed refresh
	mock.ResetServers()
	if err := s.Refresh(); err == nil {
		t.Error("Server refresh must fail")
	}
}

func TestClientServerNotFound(t *testing.T) {
	mock.ResetServers()

	cli, err := createTestClient()
	if err != nil {
		t.Error(err)
		return
	}

	s, err := cli.Server("uuid1234567")
	if s != nil {
		t.Errorf("found server %#v", s)
	}
	if err == nil {
		t.Error("error equal to nil")
	} else {
		t.Log(err)
		cs, ok := err.(*Error)
		if !ok {
			t.Error("error required to be gosigma.Error")
		}
		if cs.ServiceError.Message != "notfound" {
			t.Error("invalid error message from mock server")
		}
	}
}

func TestClientAllServersDetail(t *testing.T) {
	mock.ResetServers()

	ds := data.Server{
		ServerRecord: data.ServerRecord{
			Name:   "name",
			URI:    "uri",
			Status: "status",
			UUID:   "uuid",
		},
		Meta: map[string]string{"key1": "value1", "key2": "value2"},
	}
	mock.AddServer(&ds)

	cli, err := createTestClient()
	if err != nil {
		t.Error(err)
		return
	}

	ss, err := cli.AllServers(false)
	if err != nil {
		t.Error(err)
		return
	}

	if v, ok := ss[0].Get("key1"); ok || len(v) > 0 {
		t.Error("Error getting short server list")
	}

	ss, err = cli.AllServers(true)
	if err != nil {
		t.Error(err)
	}

	if v, ok := ss[0].Get("key1"); !ok || len(v) == 0 {
		t.Error("Error getting detailed server list")
	}

	mock.ResetServers()
}
