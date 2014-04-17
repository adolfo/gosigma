// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"github.com/Altoros/gosigma/mock"
)

func init() {
	mock.Start()
}

func TestClientQuery(t *testing.T) {

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

	ii, err := cli.Instances()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ii)
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

	ii, err := cli.Instance(*uuid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ii)
}
