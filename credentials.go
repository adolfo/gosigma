// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"net/http"
)

// Authtype defines authentication type for CloudSigma client connection
type Authtype int

const (
	// AuthtypeBasic defins HTTP basic authentication type
	AuthtypeBasic Authtype = iota
	// AuthtypeDigest defins HTTP digest authentication type
	AuthtypeDigest
	// AuthtypeCookie defins cookie authentication schema (CloudSigma specific)
	AuthtypeCookie
)

var errAuthTypeNotSupported = errors.New("authentication type is not supported")
var errEmptyUsername = errors.New("username is not allowed to be empty")
var errEmptyPassword = errors.New("password is not allowed to be empty")

// A Credentials is used to initialize new CloudSigma client
type Credentials struct {
	Type     Authtype // Authentication type
	User     string   // Username of CloudSigma account
	Password string   // Password of CloudSigma account
}

// Apply authentication credentials to HTTP request object
func (c Credentials) Apply(req *http.Request) error {
	switch c.Type {
	case AuthtypeBasic:
		req.SetBasicAuth(c.User, c.Password)
	case AuthtypeDigest, AuthtypeCookie:
		return errAuthTypeNotSupported
	}
	return nil
}

// Verify authentication credentials are valid
func (c Credentials) Verify() error {
	switch c.Type {
	case AuthtypeBasic:
	case AuthtypeDigest, AuthtypeCookie:
		return errAuthTypeNotSupported
	}

	if len(c.User) == 0 {
		return errEmptyUsername
	}

	if len(c.Password) == 0 {
		return errEmptyPassword
	}

	return nil
}
