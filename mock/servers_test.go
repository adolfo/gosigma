// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import "runtime"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Start()
}
