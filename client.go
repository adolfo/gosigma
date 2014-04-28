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
	endpoint         string
	https            *https.Client
	operationTimeout time.Duration
}

var errEmptyUsername = errors.New("username is not allowed to be empty")
var errEmptyPassword = errors.New("password is not allowed to be empty")
var errEmptyUUID = errors.New("password is not allowed to be empty")

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

// GetConnectTimeout returns connection timeout for the object
func (c Client) GetConnectTimeout() time.Duration {
	return c.https.GetConnectTimeout()
}

// ReadWriteTimeout sets read-write timeout
func (c Client) ReadWriteTimeout(timeout time.Duration) {
	c.https.ReadWriteTimeout(timeout)
}

// GetReadWriteTimeout returns connection timeout for the object
func (c Client) GetReadWriteTimeout() time.Duration {
	return c.https.GetReadWriteTimeout()
}

// OperationTimeout sets timeout for cloud operations (like cloning, starting, stopping etc)
func (c *Client) OperationTimeout(timeout time.Duration) {
	c.operationTimeout = timeout
}

// GetOperationTimeout gets timeout for cloud operations (like cloning, starting, stopping etc)
func (c Client) GetOperationTimeout() time.Duration {
	return c.operationTimeout
}

// Logger sets logger for http traces
func (c Client) Logger(logger https.Logger) {
	c.https.Logger(logger)
}

// Servers in current account
func (c Client) Servers(detail bool) ([]Server, error) {
	objs, err := c.getServers(detail)
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
func (c Client) Server(uuid string) (Server, error) {
	obj, err := c.getServer(uuid)
	if err != nil {
		return Server{}, err
	}

	srv := Server{
		client: &c,
		obj:    obj,
	}

	return srv, nil
}

// CreateServer in CloudSigma user account
func (c Client) CreateServer(components Components) (Server, error) {
	objs, err := c.createServer(components)
	if err != nil {
		return Server{}, err
	}

	if len(objs) == 0 {
		return Server{}, errors.New("no servers in response from endpoint")
	}

	s := Server{
		client: &c,
		obj:    &objs[0],
	}

	return s, nil
}

// StartServer by uuid of server instance.
func (c Client) StartServer(uuid string, avoid []string) error {
	return c.startServer(uuid, avoid)
}

// StopServer by uuid of server instance
func (c Client) StopServer(uuid string) error {
	return c.stopServer(uuid)
}

// RemoveServer by uuid of server instance with an option recursively removing attached drives.
// See RecurseXXX constants in server.go file.
func (c Client) RemoveServer(uuid, recurse string) error {
	return c.removeServer(uuid, recurse)
}

// Drives returns list of drives
func (c Client) Drives(detail, library bool) ([]Drive, error) {
	objs, err := c.getDrives(detail, library)
	if err != nil {
		return nil, err
	}

	drives := make([]Drive, len(objs))
	for i := 0; i < len(objs); i++ {
		drives[i] = Drive{
			client:  &c,
			obj:     &objs[i],
			library: library,
		}
	}

	return drives, nil
}

// Drive returns given drive by uuid
func (c Client) Drive(uuid string, library bool) (Drive, error) {
	obj, err := c.getDrive(uuid, library)
	if err != nil {
		return Drive{}, err
	}

	drv := Drive{
		client:  &c,
		obj:     obj,
		library: library,
	}

	return drv, nil
}

// CloneDrive clones given drive by uuid
func (c Client) CloneDrive(uuid string, library bool, params CloneParams, avoid []string) (Drive, error) {
	objs, err := c.cloneDrive(uuid, library, params, avoid)

	if err != nil {
		return Drive{}, err
	}

	if len(objs) == 0 {
		return Drive{}, errors.New("No object was returned from server")
	}

	drv := Drive{
		client:  &c,
		obj:     &objs[0],
		library: library,
	}

	return drv, nil
}

// Job returns job object by uuid
func (c Client) Job(uuid string) (Job, error) {
	obj, err := c.getJob(uuid)
	if err != nil {
		return Job{}, err
	}

	job := Job{
		client: &c,
		obj:    obj,
	}

	return job, nil
}

func (c Client) getServers(detail bool) ([]data.Server, error) {
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
		return nil, NewError(r, err)
	}

	return data.ReadServers(r.Body)
}

func (c Client) getServer(uuid string) (*data.Server, error) {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return nil, errEmptyUUID
	}

	u := c.endpoint + "servers/" + uuid + "/"

	r, err := c.https.Get(u, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadServer(r.Body)
}

func (c Client) startServer(uuid string, avoid []string) error {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return errEmptyUUID
	}

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
		return NewError(r, err)
	}

	return nil
}

func (c Client) stopServer(uuid string) error {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return errEmptyUUID
	}

	u := c.endpoint + "servers/" + uuid + "/action/"

	var params = make(url.Values)
	params["do"] = []string{"stop"}

	r, err := c.https.Post(u, params, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(202); err != nil {
		return NewError(r, err)
	}

	return nil
}

func (c Client) removeServer(uuid, recurse string) error {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return errEmptyUUID
	}

	u := c.endpoint + "servers/" + uuid + "/"

	var qq url.Values
	recurse = strings.TrimSpace(recurse)
	if recurse != "" {
		qq = make(url.Values)
		qq["recurse"] = []string{recurse}
	}

	r, err := c.https.Delete(u, qq, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := r.VerifyCode(204); err != nil {
		return NewError(r, err)
	}

	return nil
}

func (c Client) createServer(components Components) ([]data.Server, error) {
	// serialize
	rr, err := components.marshal()
	if err != nil {
		return nil, err
	}

	// run request
	u := c.endpoint + "servers/"
	r, err := c.https.Post(u, nil, rr)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(201); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadServers(r.Body)
}

func (c Client) getDrives(detail, library bool) ([]data.Drive, error) {
	u := c.endpoint
	if library {
		u += "libdrives"
	} else {
		u += "drives"
	}
	if detail {
		u += "/detail"
	}

	r, err := c.https.Get(u, url.Values{"limit": {"0"}})
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadDrives(r.Body)
}

func (c Client) getDrive(uuid string, library bool) (*data.Drive, error) {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return nil, errEmptyUUID
	}

	u := c.endpoint
	if library {
		u += "libdrives/"
	} else {
		u += "drives/"
	}
	u += uuid + "/"

	r, err := c.https.Get(u, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadDrive(r.Body)
}

func (c Client) cloneDrive(uuid string, library bool, params CloneParams, avoid []string) ([]data.Drive, error) {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return nil, errEmptyUUID
	}

	u := c.endpoint
	if library {
		u += "libdrives/"
	} else {
		u += "drives/"
	}
	u += uuid + "/action/"

	var qq = make(url.Values)
	qq["do"] = []string{"clone"}

	if len(avoid) > 0 {
		qq["avoid"] = []string{strings.Join(avoid, ",")}
	}

	rr, err := params.makeJsonReader()
	if err != nil {
		return nil, err
	}

	r, err := c.https.Post(u, qq, rr)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(202); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadDrives(r.Body)
}

func (c Client) getJob(uuid string) (*data.Job, error) {
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		return nil, errEmptyUUID
	}

	u := c.endpoint + "jobs/" + uuid + "/"

	r, err := c.https.Get(u, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := r.VerifyJSON(200); err != nil {
		return nil, NewError(r, err)
	}

	return data.ReadJob(r.Body)
}
