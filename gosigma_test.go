// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "testing"

func TestVersionStringMatches(t *testing.T) {
	if vs, vns := Version(), VersionNumber().String(); vs != vns {
		t.Errorf("Version() != VersionNumber().String(): '%s' != '%s'")
	}
}

func TestGetRegionEndpoint(t *testing.T) {
	check := func(r string, url string) {
		ep, err := GetRegionEndpoint(r)
		if err != nil {
			t.Error(err)
		}
		if ep != url {
			t.Errorf("ep value = '%s', wants '%s'", ep, url)
		}
	}

	check("zrh", "https://zrh.cloudsigma.com/api/2.0/")
	check("lvs", "https://lvs.cloudsigma.com/api/2.0/")
}
