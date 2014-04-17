// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"github.com/Altoros/gosigma/data"
	"github.com/Altoros/gosigma/mock"
)

var mockEndpoint string

func init() {
	mock.Start()
	mockEndpoint = mock.Endpoint("")
}

func createClient() (*Client, error) {
	return NewClient(mockEndpoint, mock.TestUser, mock.TestPassword, nil)
}

func TestClientCreation(t *testing.T) {
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
	cli, err := createClient()
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
	}
}

func TestAllServersEmpty(t *testing.T) {
	cli, err := createClient()
	if err != nil || cli == nil {
		t.Error("NewClient() failed:", err, cli)
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

func TestAllServers(t *testing.T) {
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

	cli, err := createClient()
	if err != nil {
		t.Error(err)
	}

	servers, err := cli.AllServers(true)
	if err != nil {
		t.Error(err)
	}

	if len(servers) != 1 {
		t.Errorf("invalid len: %v", servers)
	}

	s := servers[0]

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
	}

	mock.ResetServers()
}

func TestServer(t *testing.T) {
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

	cli, err := createClient()
	if err != nil {
		t.Error(err)
	}

	if s, err := cli.Server(""); err == nil || s != nil {
		t.Error(err)
	}

	s, err := cli.Server("uuid")
	if err != nil || s == nil {
		t.Error(err)
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

	mock.ResetServers()
}

var cloud = flag.String("cloud", "", "run tests against CloudSigma endpoint, specify credentials in form user:pass as parameter")
var uuid = flag.String("uuid", "", "uuid of server at CloudSigma to run server specific tests")

func getCloudCredentials() ([]string, error) {
	if cloud == nil || *cloud == "" {
		return nil, nil
	}
	parts := strings.SplitN(*cloud, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("Invalid credentials: " + *cloud)
	}
	if len(parts[0]) == 0 {
		return nil, errors.New("Invalid credentials: " + *cloud)
	}
	return parts, nil
}

func TestCloudServers(t *testing.T) {
	cr, err := getCloudCredentials()
	if cr == nil {
		if err == nil {
			t.SkipNow()
		} else {
			t.Error(err)
		}
		return
	}

	cli, err := NewClient(DefaultRegion, cr[0], cr[1], nil)
	if err != nil {
		t.Error(err)
	}

	ii, err := cli.AllServers(false)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", ii)
}

func TestCloudServer(t *testing.T) {
	cr, err := getCloudCredentials()
	if cr == nil {
		if err == nil {
			t.SkipNow()
		} else {
			t.Error(err)
		}
		return
	}

	cli, err := NewClient(DefaultRegion, cr[0], cr[1], nil)
	if err != nil {
		t.Error(err)
	}

	ii, err := cli.Server(*uuid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", ii)
}
