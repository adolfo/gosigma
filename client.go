// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"crypto/tls"
	"errors"
	"net/url"

	"github.com/Altoros/gosigma/data"
	"github.com/Altoros/gosigma/https"
)

// A Client sends and receives requests to CloudSigma endpoint
type Client struct {
	endpoint string
	https    *https.Client
}

var errEmptyUsername = errors.New("username is not allowed to be empty")
var errEmptyPassword = errors.New("password is not allowed to be empty")

// NewClient returns new CloudSigma client object
func NewClient(endpoint string,
	username, password string, tlsConfig *tls.Config) (*Client, error) {

	endpoint, err := ResolveEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	if len(username) == 0 {
		return nil, errEmptyUsername
	}

	if len(password) == 0 {
		return nil, errEmptyPassword
	}

	client := &Client{
		endpoint: endpoint,
		https:    https.NewAuthClient(username, password, tlsConfig),
	}

	return client, nil
}

// Instances of all servers in current account
func (c Client) Instances() ([]data.Server, error) {
	u := c.endpoint + "servers"
	r, err := c.https.GetQuery(u, url.Values{"limit": {"0"}})
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return data.ReadServers(r.Body)
}

// Instance description for given server uuid
func (c Client) Instance(uuid string) (*data.Server, error) {
	u := c.endpoint + "servers/" + uuid
	r, err := c.https.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return data.ReadServer(r.Body)
}
