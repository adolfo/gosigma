// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"bytes"
	"testing"
)

func init() {
	Start()
}

func testEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestServerMockSections(t *testing.T) {

	ch := make(chan int)

	check := func(s string) {
		resp, err := Request(s)
		if err != nil {
			t.Errorf("Section %s: %s", s, err)
		}
		defer resp.Body.Close()

		id := GetIDFromResponse(resp)
		jj := GetJournal(id)
		Log(t, jj)

		if resp.StatusCode != 200 {
			t.Errorf("Section %s: %s", s, resp.Status)
		}

		if len(jj) == 0 {
			t.Errorf("Section %s: journal length is zero")
		}

		j := jj[0]

		var buf bytes.Buffer
		buf.ReadFrom(resp.Body)

		if !testEq(j.Response.Body.Bytes(), buf.Bytes()) {
			t.Errorf("Section: body error")
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
