// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package comm

import (
	"crypto/tls"
	"errors"
	"net/http"
)

// NewHttpsClient returns http.Client object with configured https transport
func NewHttpsClient(tlsConfig *tls.Config) *http.Client {
	if tlsConfig == nil {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	redirectChecker := func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		lastReq := via[len(via)-1]
		if auth := lastReq.Header.Get("Authorization"); len(auth) > 0 {
			req.Header.Add("Authorization", auth)
		}
		return nil
	}

	https := &http.Client{
		Transport:     tr,
		CheckRedirect: redirectChecker,
	}

	return https
}
