// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import "testing"

func init() {
	Start()
}

func TestServerMockSections(t *testing.T) {

	ch := make(chan int)

	check := func(s string) {
		resp, err := Request(s)
		if err != nil {
			t.Errorf("Section %s: %s", s, err)
		}

		id := GetIDFromResponse(resp)
		j := GetJournal(id)
		Log(t, j)

		if resp.StatusCode != 200 {
			t.Errorf("Section %s: %s", s, resp.Status)
		}

		ch <- 1
	}

	const sectionCount = 2
	go check("capabilities")
	go check("drives")

	var s int = 0
	for s < sectionCount {
		s += <-ch
	}
}
