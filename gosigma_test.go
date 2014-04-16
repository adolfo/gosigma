// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "testing"

func TestVersionStringMatches(t *testing.T) {
	t.Parallel()
	if vs, vns := Version(), VersionNumber().String(); vs != vns {
		t.Errorf("Version() != VersionNumber().String(): '%s' != '%s'")
	}
}
