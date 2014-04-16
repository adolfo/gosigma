// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
)

type Authtype int

const (
	AuthtypeBasic Authtype = iota
	AuthtypeDigest
	AuthtypeCookie
)

var errAuthTypeNotSupported error = errors.New("authentication type is not supported")
var errEmptyEndpoint error = errors.New("endpoint are not allowed to be empty")
var errHttpsRequired error = errors.New("endpoint must use https scheme")
var errInvalidAuth error = errors.New("auth information is not allowed in the endpoint string")
var errEndpointWithQuery error = errors.New("query information is not allowed in the endpoint string")
var errEmptyUsername error = errors.New("username is not allowed to be empty")
var errEmptyPassword error = errors.New("password is not allowed to be empty")

// A Credentials is used to initialize new CloudSigma client
type Credentials struct {
	Type     Authtype // Authentication type
	User     string   // Username of CloudSigma account
	Password string   // Password of CloudSigma account
}

func (c *Credentials) Set(req *http.Request) error {
	switch c.Type {
	case AuthtypeBasic:
		req.SetBasicAuth(c.User, c.Password)
	case AuthtypeDigest, AuthtypeCookie:
		return errAuthTypeNotSupported
	}
	return nil
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
	https       *http.Client
}

// NewClient returns new CloudSigma client object
func NewClient(c Configuration, tlsConfig *tls.Config) (*Client, error) {
	if len(c.Endpoint) == 0 {
		return nil, errEmptyEndpoint
	}

	endpoint, err := GetRegionEndpoint(c.Endpoint)
	if err != nil {
		endpoint = c.Endpoint
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		return nil, errHttpsRequired
	}
	if u.User != nil {
		return nil, errInvalidAuth
	}
	if len(u.RawQuery) > 0 || len(u.Fragment) > 0 {
		return nil, errEndpointWithQuery
	}

	if len(c.User) == 0 {
		return nil, errEmptyUsername
	}

	if len(c.Password) == 0 {
		return nil, errEmptyPassword
	}

	switch c.Type {
	case AuthtypeBasic:
	case AuthtypeDigest, AuthtypeCookie:
		return nil, errAuthTypeNotSupported
	}

	client := &Client{
		endpoint:    endpoint,
		credentials: c.Credentials,
		https:       NewHttpsClient(tlsConfig),
	}

	return client, nil
}

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

func (c *Client) Endpoint() string {
	return c.endpoint
}

func (c *Client) Instances() (ii []Instance, err error) {
	return
}

func (c *Client) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	err = c.credentials.Set(req)
	if err != nil {
		return nil, err
	}

	return c.https.Do(req)
}
