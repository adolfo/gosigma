// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "testing"

func TestSetupShortEndpoints(t *testing.T) {

	check := func(ep string, url string) {
		c := Configuration{Credentials{AuthtypeBasic, "user", "pwd"}, ep}

		client, err := NewClient(c)
		if err != nil {
			t.Error(err)
		}

		if ep := client.Endpoint(); ep != url {
			t.Errorf("ep value = '%s', wants '%s'", ep, url)
		}
	}

	check("zrh", "https://zrh.cloudsigma.com/api/2.0/")
	check("lvs", "https://lvs.cloudsigma.com/api/2.0/")
}
