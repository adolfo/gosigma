// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

// A ContextNIC interface represents network interface card for server instance context
type ContextNIC interface {
	fmt.Stringer
	Mac() string
	Model() string
}

// A ContextIPv4 interface represents IPv4 information for server instance context
type ContextIPv4 interface {
	fmt.Stringer
	Gateway() string
	Nameservers() []string
	Netmask() string
	UUID() string
}

// A ContextVLan interface represents VLan information for server instance context
type ContextVLan interface {
	fmt.Stringer
	UUID() string
}

///////////////////////////////////////////////////////////////////////////////////

// A context implements network interface card for server instance context
type contextNIC struct {
	obj *data.ContextNIC
}

var _ ContextNIC = contextNIC{}

func (c contextNIC) Mac() string   { return c.obj.Mac }
func (c contextNIC) Model() string { return c.obj.Model }
func (c contextNIC) IPv4() ContextIPv4 {
	if c.obj.IPv4 != nil {
		return contextIPv4{c.obj.IPv4}
	} else {
		return nil
	}
}
func (c contextNIC) VLAN() ContextVLan {
	if c.obj.VLan != nil {
		return contextVLan{c.obj.VLan}
	} else {
		return nil
	}
}

func (c contextNIC) String() string {
	return fmt.Sprintf("{Mac: %q\nModel: %q}", c.Mac(), c.Model())
}

///////////////////////////////////////////////////////////////////////////////////

// A contextIPv4 implements IPv4 information for server instance context
type contextIPv4 struct {
	obj *data.ContextIPv4
}

var _ ContextIPv4 = contextIPv4{}

func (ci contextIPv4) Gateway() string       { return ci.obj.IP.Gateway }
func (ci contextIPv4) Nameservers() []string { return ci.obj.IP.Nameservers }
func (ci contextIPv4) Netmask() string       { return ci.obj.IP.Netmask }
func (ci contextIPv4) UUID() string          { return ci.obj.IP.UUID }

func (ci contextIPv4) String() string {
	return fmt.Sprintf("{Gateway: %q, Netmask: %q, UUID: %q}",
		ci.Gateway(), ci.Netmask(), ci.UUID())
}

///////////////////////////////////////////////////////////////////////////////////

// A contextVLan implements VLan information for server instance context
type contextVLan struct {
	obj *data.ContextVLan
}

var _ ContextVLan = contextVLan{}

func (cv contextVLan) UUID() string { return cv.obj.UUID }

func (cv contextVLan) String() string {
	return fmt.Sprintf("{UUID: %q}", cv.UUID())
}
