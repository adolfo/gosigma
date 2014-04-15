// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"net/url"
)

type Authtype int

const (
	AuthtypeBasic Authtype = iota
	AuthtypeDigest
	AuthtypeCookie
)

// A Credentials is used to initialize new CloudSigma client
type Credentials struct {
	Type     Authtype // Authentication type
	User     string   // CloudSigma account username
	Password string   // CloudSigma account password
}

// A Configuration is used to initialize new CloudSigma client
type Configuration struct {
	Credentials
	Endpoint string // Endpoint short name (zrh,lvs,...) or https URL
}

// A Client sends and receives requests to CloudSigma endpoint
type Client struct {
	endpoint    string
	credentials Credentials
}

// NewClient returns new CloudSigma client object
func NewClient(c Configuration) (*Client, error) {
	if len(c.Endpoint) == 0 {
		return nil, errors.New("endpoint are not allowed to be empty")
	}

	endpoint := c.Endpoint
	switch endpoint {
	case "zrh":
		endpoint = "https://zrh.cloudsigma.com/api/2.0/"
	case "lvs":
		endpoint = "https://lvs.cloudsigma.com/api/2.0/"
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		return nil, errors.New("endpoint must use https scheme")
	}
	if u.User != nil {
		return nil, errors.New("user information is not allowed in the endpoint string")
	}
	if len(u.RawQuery) > 0 || len(u.Fragment) > 0 {
		return nil, errors.New("query information is not allowed in the endpoint string")
	}

	if len(c.User) == 0 {
		return nil, errors.New("username are not allowed to be empty")
	}

	if len(c.Password) == 0 {
		return nil, errors.New("password are not allowed to be empty")
	}

	switch c.Type {
	case AuthtypeBasic:
	case AuthtypeDigest, AuthtypeCookie:
		return nil, errors.New("authentication type is not supported now")
	}

	client := &Client{
		endpoint:    endpoint,
		credentials: c.Credentials,
	}

	return client, nil
}

func (c Client) Endpoint() string {
	return c.endpoint
}
