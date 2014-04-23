// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDataServersReaderFail(t *testing.T) {
	r := failReader{}

	if _, err := ReadServer(r); err == nil || err.Error() != "test error" {
		t.Error("Fail")
	}

	if _, err := ReadServers(r); err == nil || err.Error() != "test error" {
		t.Error("Fail")
	}
}

func TestDataServersUnmarshal(t *testing.T) {
	var ss Servers
	ss.Meta.Limit = 12345
	ss.Meta.Offset = 12345
	ss.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonServersData), &ss)
	if err != nil {
		t.Error(err)
	}

	verifyMeta(t, &ss.Meta, 0, 0, 5)

	for i := 0; i < len(serversData); i++ {
		compareServers(t, i, &ss.Objects[i], &serversData[i])
	}
}

func TestDataServersReadServers(t *testing.T) {
	ss, err := ReadServers(strings.NewReader(jsonServersData))
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(serversData); i++ {
		compareServers(t, i, &ss[i], &serversData[i])
	}
}

func TestDataServersDetailUnmarshal(t *testing.T) {
	var ss Servers
	ss.Meta.Limit = 12345
	ss.Meta.Offset = 12345
	ss.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonServersDetailData), &ss)
	if err != nil {
		t.Error(err)
	}

	verifyMeta(t, &ss.Meta, 0, 0, 5)

	for i := 0; i < len(serversData); i++ {
		compareServers(t, i, &ss.Objects[i], &serversDetailData[i])
	}
}

func TestDataServersReadServersDetail(t *testing.T) {
	ss, err := ReadServers(strings.NewReader(jsonServersDetailData))
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(serversData); i++ {
		compareServers(t, i, &ss[i], &serversDetailData[i])
	}
}

func TestDataServersReadServerDetail(t *testing.T) {
	s, err := ReadServer(strings.NewReader(jsonServerData))
	if err != nil {
		t.Error(err)
	}
	compareServers(t, 0, s, &serverData)
}

func compareServers(t *testing.T, i int, value, wants *Server) {
	if value.ServerRecord != wants.ServerRecord {
		t.Errorf("ServerRecord error [%d]: found %#v, wants %#v", i, value.ServerRecord, wants.ServerRecord)
	}
	if value.CPU != wants.CPU {
		t.Errorf("Server.CPU error [%d]: found %#v, wants %#v", i, value.CPU, wants.CPU)
	}
	if value.Mem != wants.Mem {
		t.Errorf("Server.Mem error [%d]: found %#v, wants %#v", i, value.Mem, wants.Mem)
	}
	if len(value.Meta) != len(wants.Meta) {
		t.Errorf("Server.Meta error [%d]: found %#v, wants %#v", i, value.Meta, wants.Meta)
	}
	if len(value.NICs) != len(wants.NICs) {
		t.Errorf("Server.NICs error [%d]: found %#v, wants %#v", i, value.NICs, wants.NICs)
	}
	if len(value.Drives) != len(wants.Drives) {
		t.Errorf("Server.Drives error [%d]: found %#v, wants %#v", i, value.Drives, wants.Drives)
	}
}
