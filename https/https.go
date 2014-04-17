// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package https

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
)

// Client represents HTTPS client connection with optional basic authentication
type Client struct {
	protocol *http.Client
	username string
	password string
}

// NewClient returns new Client object with configured https transport
func NewClient(tlsConfig *tls.Config) *Client {
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

	https := &Client{
		protocol: &http.Client{
			Transport:     tr,
			CheckRedirect: redirectChecker,
		},
	}

	return https
}

// NewAuthClient returns new Client object with configured https transport
// and attached authentication
func NewAuthClient(username, password string, tlsConfig *tls.Config) *Client {
	https := NewClient(tlsConfig)
	https.username = username
	https.password = password
	return https
}

func (c Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(c.username) != 0 {
		req.SetBasicAuth(c.username, c.password)
	}

	return c.protocol.Do(req)
}

func (c Client) GetQuery(url string, values url.Values) (*http.Response, error) {
	if len(values) == 0 {
		return c.Get(url)
	}

	url += "?" + values.Encode()

	return c.Get(url)
}
