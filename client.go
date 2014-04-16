// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/Altoros/gosigma/comm"
	"github.com/Altoros/gosigma/data"
)

// A Client sends and receives requests to CloudSigma endpoint
type Client struct {
	endpoint string
	cred     Credentials
	https    *http.Client
}

// NewClient returns new CloudSigma client object
func NewClient(endpoint string, credentials Credentials,
	tlsConfig *tls.Config) (*Client, error) {

	// check the endpoint is a region name
	resolved, err := GetRegionEndpoint(endpoint)
	if err == nil {
		endpoint = resolved
	}

	err = VerifyEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	err = credentials.Verify()
	if err != nil {
		return nil, err
	}

	client := &Client{
		endpoint: endpoint,
		cred:     credentials,
		https:    comm.NewHttpsClient(tlsConfig),
	}

	return client, nil
}

func (c *Client) Endpoint() string {
	return c.endpoint
}

func (c *Client) Instances() (ii []data.Server, err error) {
	return
}

func (c *Client) get(url string) (*http.Response, error) {
	url = c.Endpoint() + url

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	err = c.cred.Apply(req)
	if err != nil {
		return nil, err
	}

	return c.https.Do(req)
}

func (c *Client) query(query string, values url.Values) (*http.Response, error) {
	if len(values) == 0 {
		return c.get(query)
	}

	query += "?" + values.Encode()

	return c.get(query)
}
