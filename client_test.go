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

var cloud *string = flag.String("cloud", "", "run tests against CloudSigma endpoint, specify credentials in form user:pass as parameter")

func getCloudCredentials() (*Credentials, error) {
	if cloud == nil || *cloud == "" {
		return nil, nil
	}
	parts := strings.SplitN(*cloud, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("Invalid credentials: " + *cloud)
	}
	cr := Credentials{AuthtypeBasic, parts[0], parts[1]}
	if err := cr.Verify(); err != nil {
		return nil, errors.New("Invalid credentials: " + *cloud)
	}
	return &cr, nil
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

	cli, err := NewClient(DefaultRegion, *cr, nil)
	if err != nil {
		t.Error(err)
	}

	ii, err := cli.Instances()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ii)
}
