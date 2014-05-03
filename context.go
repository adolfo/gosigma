// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

// A Context interface represents server instance context in CloudSigma account
type Context interface {
	fmt.Stringer
	Cpu() int64
	Mem() int64
	Name() string
	NICs() []ContextNIC
	UUID() string
	VNCPassword() string
	Get(key string) (string, bool)
}

// A context implements server instance context in CloudSigma account
type context struct {
	obj *data.Context
}

var _ Context = context{}

// Cpu frequency in MHz
func (c context) Cpu() int64 { return c.obj.CPU }

// Mem capacity in bytes
func (c context) Mem() int64 { return c.obj.Mem }

// Name of server instance
func (c context) Name() string { return c.obj.Name }

// NICs for this context instance.
func (c context) NICs() []ContextNIC {
	r := make([]ContextNIC, 0, len(c.obj.NICs))
	for i := range c.obj.NICs {
		nic := contextNIC{&c.obj.NICs[i]}
		r = append(r, nic)
	}
	return r
}

// UUID of server instance
func (c context) UUID() string { return c.obj.UUID }

// VNCPassword to access the server
func (c context) VNCPassword() string { return c.obj.VNCPassword }

// Get meta-information value stored in the server instance
func (c context) Get(key string) (v string, ok bool) {
	v, ok = c.obj.Meta[key]
	return
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (c context) String() string {
	return fmt.Sprintf("{Name: %q\nUUID: %q}", c.Name(), c.UUID())
}
