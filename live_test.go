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
var suid = flag.String("suid", "", "uuid of server at CloudSigma to run server specific tests")
var duid = flag.String("duid", "", "uuid of drive at CloudSigma to run drive specific tests")
var force = flag.String("force", "n", "force start/stop live tests")

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

	if *trace != "n" {
		cli.Logger(t)
	}

	ii, err := cli.Servers(false)
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

	if *suid == "" {
		t.Skip("-suid=<server-uuid> must be specified")
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	s, err := cli.Server(*suid)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", s)
}

func TestLiveStart(t *testing.T) {
	u, p, err := parseCredentials()
	if u == "" {
		skipTest(t, err)
		return
	}

	if *suid == "" {
		t.Skip("-suid=<server-uuid> must be specified")
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	s, err := cli.Server(*suid)
	if err != nil {
		t.Error(err)
		return
	}

	if s.Status() != ServerStopped && *force == "n" {
		t.Skip("wrong server status", s.Status())
		return
	}

	if err := s.Start(); err != nil {
		t.Error(err)
	}
}

func TestLiveStop(t *testing.T) {
	u, p, err := parseCredentials()
	if u == "" {
		skipTest(t, err)
		return
	}

	if *suid == "" {
		t.Skip("-suid=<server-uuid> must be specified")
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	s, err := cli.Server(*suid)
	if err != nil {
		t.Error(err)
		return
	}

	if s.Status() != ServerRunning && *force == "n" {
		t.Skip("wrong server status", s.Status())
		return
	}

	if err := s.Stop(); err != nil {
		t.Error(err)
	}
}

func TestLiveDrive(t *testing.T) {
	u, p, err := parseCredentials()
	if u == "" {
		skipTest(t, err)
		return
	}

	if *duid == "" {
		t.Skip("-duid=<drive-uuid> must be specified")
		return
	}

	cli, err := NewClient(DefaultRegion, u, p, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if *trace != "n" {
		cli.Logger(t)
	}

	d, err := cli.Drive(*duid)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", d)
}
