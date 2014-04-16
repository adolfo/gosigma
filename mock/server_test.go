// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"bytes"
	"net/http"
	"runtime"
	"testing"
	"github.com/Altoros/gosigma/comm"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Start()
}

func TestAuth(t *testing.T) {

	t.Parallel()

	check := func(u, p string, wants int) {
		req, err := http.NewRequest("GET", Endpoint()+"capabilities", nil)
		if err != nil {
			t.Error(err)
		}
		if u != "" {
			req.SetBasicAuth(u, p)
		}
		client := comm.NewHttpsClient(nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != wants {
			t.Errorf("Status = %d, wants %d", resp.StatusCode, wants)
		}
	}

	go check("", "", 401)
	go check(TestUser, "", 401)
	go check(TestUser, "1", 401)
	go check(TestUser+"1", TestPassword, 401)
	go check(TestUser, TestPassword+"1", 401)
	go check(TestUser, TestPassword, 200)
}

func TestSections(t *testing.T) {

	t.Parallel()

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

		if !bytes.Equal(j.Response.Body.Bytes(), buf.Bytes()) {
			t.Errorf("Section: body error")
		}

		ch <- 1
	}

	const sectionCount = 3
	go check("capabilities")
	go check("drives")
	go check("servers")

	var s int = 0
	for s < sectionCount {
		s += <-ch
	}
}

/*
func TestHeaders(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := "https://localhost/headers.php"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}

	req.SetBasicAuth("test@example.com", "test")
	s := req.Header.Get("Authorization")
	t.Log("Auth:", s)

	resp, _ := client.Do(req)
	body, err := httputil.DumpResponse(resp, true)
	t.Log(string(body))
}
*/
