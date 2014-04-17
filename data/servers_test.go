// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"strings"
	"testing"
	"testing/iotest"
)

func verifyServerObject(t *testing.T, i int, s Server, name, uri, status, uuid string) {
	if s.Name != name {
		t.Errorf("Object %d, Name = '%s', wants '%s'", i, s.Name, name)
	}
	if s.URI != uri {
		t.Errorf("Object %d, URI = '%s', wants '%s'", i, s.URI, uri)
	}
	if s.Status != status {
		t.Errorf("Object %d, Status = '%s', wants '%s'", i, s.Status, status)
	}
	if s.UUID != uuid {
		t.Errorf("Object %d, UUID = '%s', wants '%s'", i, s.UUID, uuid)
	}
}

func verifyServerObjects(t *testing.T, ii []Server) {
	if len(ii) != 5 {
		t.Errorf("Meta.Objects.len = %d, wants 5", len(ii))
	}

	verify := func(i int, name, uri, status, uuid string) {
		verifyServerObject(t, i, ii[i], name, uri, status, uuid)
	}

	verify(0, "test_server_4", "/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
		"stopped", "43b1110a-31c5-41cc-a3e7-0b806076a913")
	verify(1, "test_server_2", "/api/2.0/servers/3be1ebc6-1d03-4c4b-88ff-02557b940d19/",
		"stopped", "3be1ebc6-1d03-4c4b-88ff-02557b940d19")
	verify(2, "test_server_0", "/api/2.0/servers/b1defe23-e725-474d-acba-e46baa232611/",
		"stopped", "b1defe23-e725-474d-acba-e46baa232611")
	verify(3, "test_server_3", "/api/2.0/servers/cff0f338-2b84-4846-a028-3ec9e1b86184/",
		"stopped", "cff0f338-2b84-4846-a028-3ec9e1b86184")
	verify(4, "test_server_1", "/api/2.0/servers/93a04cd5-84cb-41fc-af17-683e3868ee95/",
		"stopped", "93a04cd5-84cb-41fc-af17-683e3868ee95")
}

func verifyMeta(t *testing.T, m *Meta, limit, offset, count int) {
	if m.Limit != limit {
		t.Errorf("Meta.Limit = %d, wants %d", m.Limit, limit)
	}
	if m.Offset != offset {
		t.Errorf("Meta.Offset = %d, wants %d", m.Offset, offset)
	}
	if m.TotalCount != count {
		t.Errorf("Meta.TotalCount = %d, wants %d", m.TotalCount, count)
	}
}

func verifyServers(t *testing.T, ii *Servers) {
	verifyMeta(t, &ii.Meta, 0, 0, 5)
	verifyServerObjects(t, ii.Objects)
}

func TestUnmarshal(t *testing.T) {
	var ii Servers
	ii.Meta.Limit = 12345
	ii.Meta.Offset = 12345
	ii.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonServersData), &ii)
	if err != nil {
		t.Error(err)
	}
	verifyServers(t, &ii)
}

func TestReadServers(t *testing.T) {
	servers, err := ReadServers(strings.NewReader(jsonServersData))
	if err != nil {
		t.Error(err)
	}
	verifyServerObjects(t, servers)
}

func TestReadServersHalf(t *testing.T) {
	r := strings.NewReader(jsonServersData)
	servers, err := ReadServers(iotest.HalfReader(r))
	if err != nil {
		t.Error(err)
	}
	verifyServerObjects(t, servers)
}

func verifyNIC(t *testing.T, i int, n NIC, conf, uri, uuid string) {
	if n.IPv4.Conf != conf {
		t.Errorf("nic.IPv4.Conf for (idx: %d) %+v", i, n)
	}
	if n.IPv4.IP.URI != uri {
		t.Errorf("nic.IPv4.URI for (idx: %d) %+v", i, n)
	}
	if n.IPv4.IP.UUID != uuid {
		t.Errorf("nic.IPv4.UUID for (idx: %d) %+v", i, n)
	}
}

func TestReadServersDetail(t *testing.T) {
	servers, err := ReadServers(strings.NewReader(jsonServersDetailData))
	if err != nil {
		t.Error(err)
	}
	verifyServerObjects(t, servers)

	// # verify NICs
	server := servers[0]

	verifyNIC(t, 0, server.NICs[0], "static", "/api/2.0/ips/31.171.246.37/", "31.171.246.37")
	verifyNIC(t, 1, server.NICs[1], "", "", "")
}

func TestReadServerDetail(t *testing.T) {
	server, err := ReadServer(strings.NewReader(jsonServerData))
	if err != nil {
		t.Error(err)
	}
	verifyServerObject(t, 0, *server, "trusty-server-cloudimg-amd64",
		"/api/2.0/servers/472835d5-2bbb-4d87-9d08-7364bc373691/",
		"starting", "472835d5-2bbb-4d87-9d08-7364bc373691")

	// # verify NICs
	verifyNIC(t, 0, server.NICs[0], "static", "/api/2.0/ips/31.171.246.37/", "31.171.246.37")
	verifyNIC(t, 1, server.NICs[1], "", "", "")
}
