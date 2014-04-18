// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"flag"
	"strings"
	"testing"
)

var live = flag.String("live", "", "run live tests against CloudSigma endpoint, specify credentials in form -live=user:pass")
var uuid = flag.String("uuid", "", "uuid of server at CloudSigma to run server specific tests")

func parseCredentials() (u string, p string, e error) {
	if *live == "" {
		return
	}

	parts := strings.SplitN(*live, ":", 2)
	if len(parts) != 2 || parts[0] == "" {
		e = errors.New("Invalid credentials: " + *live)
		return
	}

	u, p = parts[0], parts[1]

	return
}

func skipTest(t *testing.T, e error) {
	if e == nil {
		t.SkipNow()
	} else {
		t.Error(e)
	}
}

func TestLiveServers(t *testing.T) {
	u, p, err := parseCredentials()
	if u == "" {
		skipTest(t, err)
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	ii, err := cli.AllServers(false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", ii)
}

func TestLiveServer(t *testing.T) {
	u, p, err := parseCredentials()
	if u == "" {
		skipTest(t, err)
		return
	}

	if *uuid == "" {
		t.Skip("-uuid=<server-uuid> must be specified")
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	ii, err := cli.Server(*uuid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", ii)
}
