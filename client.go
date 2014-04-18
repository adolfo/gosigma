// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"crypto/tls"
	"errors"
	"net/url"
	"strings"
	"time"

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
func NewClient(endpoint string, username, password string,
	tlsConfig *tls.Config) (*Client, error) {

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

// ConnectTimeout sets connection timeout
func (c Client) ConnectTimeout(timeout time.Duration) {
	c.https.ConnectTimeout(timeout)
}

// ReadWriteTimeout sets read-write timeout
func (c Client) ReadWriteTimeout(timeout time.Duration) {
	c.https.ReadWriteTimeout(timeout)
}

// Logger sets logger for http traces
func (c Client) Logger(logger https.Logger) {
	c.https.Logger(logger)
}

// AllServers in current account
func (c Client) AllServers(detail bool) ([]Server, error) {
	objs, err := c.getAllServers(detail)
	if err != nil {
		return nil, err
	}

	servers := make([]Server, len(objs))
	for i := 0; i < len(objs); i++ {
		servers[i] = Server{
			client: &c,
			obj:    &objs[i],
		}
	}

	return servers, nil
}

// Server returns given server by uuid
func (c Client) Server(uuid string) (*Server, error) {
	obj, err := c.getServer(uuid)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		client: &c,
		obj:    obj,
	}

	return srv, nil
}

// StartServer by uuid
func (c Client) StartServer(uuid string, avoid []string) error {
	return c.startServer(uuid, avoid)
}

// StopServer by uuid
func (c Client) StopServer(uuid string) error {
	return c.stopServer(uuid)
}

func (c Client) getAllServers(detail bool) ([]data.Server, error) {
	u := c.endpoint + "servers"
	if detail {
		u += "/detail"
	}

	r, err := c.https.Get(u, url.Values{"limit": {"0"}})
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, err
	}

	return data.ReadServers(r.Body)
}

func (c Client) getServer(uuid string) (*data.Server, error) {
	u := c.endpoint + "servers/" + uuid

	r, err := c.https.Get(u, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, err
	}

	return data.ReadServer(r.Body)
}

func (c Client) startServer(uuid string, avoid []string) error {
	u := c.endpoint + "servers/" + uuid + "/action/"

	var params = make(url.Values)
	params["do"] = []string{"start"}

	if len(avoid) > 0 {
		params["avoid"] = []string{strings.Join(avoid, ",")}
	}

	r, err := c.https.Post(u, params, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(202); err != nil {
		return err
	}

	return nil
}

func (c Client) stopServer(uuid string) error {
	return nil
}
