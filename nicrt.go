// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

// A RuntimeNIC interface represents runtime information for network interface card
type RuntimeNIC interface {
	// Convert to string
	fmt.Stringer

	// AddressIPv4 returns runtime IPv4 address (if any)
	AddressIPv4() string

	// Type of network interface card (public, private, etc)
	Type() string
}

// A runtimeNIC implements runtime information for network interface card
type runtimeNIC struct {
	obj *data.RuntimeNetwork
}

var _ RuntimeNIC = runtimeNIC{}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (r runtimeNIC) String() string {
	return fmt.Sprintf(`{Type: %q, Address: %q}`, r.Type(), r.AddressIPv4())
}

// AddressIPv4 returns runtime IPv4 address (if any)
func (r runtimeNIC) AddressIPv4() string {
	if r.obj != nil && r.obj.IPv4 != nil {
		return r.obj.IPv4.UUID
	} else {
		return ""
	}
}

// Type of network interface card (public, private, etc)
func (r runtimeNIC) Type() string {
	if r.obj != nil {
		return r.obj.InterfaceType
	} else {
		return ""
	}
}
