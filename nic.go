// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

const (
	NIC_private = "private"
	NIC_public  = "public"
)

// A NIC interface represents network interface card instance
type NIC interface {
	// Convert to string
	fmt.Stringer

	// Address returns address configured for network interface card
	Address() string

	// Conf returns type of network interface card configuration. 'private' for NIC_vlan type,
	// 'static', 'dhcp', 'manual' for NIC_ip
	Conf() string

	// MAC address
	MAC() string

	// Model of virtual network interface card
	Model() string

	// Runtime returns runtime information for network interface card or nil if stopped
	Runtime() RuntimeNIC

	// Type of virtual network interface card (private, public)
	Type() string
}

// A nic implements network interface card instance in CloudSigma
type nic struct {
	client *Client
	obj    *data.NIC
}

var _ NIC = nic{}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (n nic) String() string {
	return fmt.Sprintf(`{Type: %q, Conf: %q, Model: %q, MAC: %q, Address: %q}`,
		n.Type(), n.Conf(), n.Model(), n.MAC(), n.Address())
}

// Address returns address configured for network interface card
func (n nic) Address() string {
	if n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return n.obj.VLAN.UUID
	}
	if n.obj.IPv4 != nil && n.obj.IPv4.IP.UUID != "" {
		return n.obj.IPv4.IP.UUID
	}
	return ""
}

// Conf returns type of network interface card configuration. 'private' for NIC_vlan type,
// 'static', 'dhcp', 'manual' for NIC_ip
func (n nic) Conf() string {
	if n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return NIC_private
	}
	if n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return n.obj.IPv4.Conf
	}
	return ""
}

// MAC address
func (n nic) MAC() string { return n.obj.MAC }

// Model of virtual network interface card
func (n nic) Model() string { return n.obj.Model }

// Runtime returns runtime information for network interface card or nil if stopped
func (n nic) Runtime() RuntimeNIC {
	if n.obj.Runtime != nil {
		return runtimeNIC{n.obj.Runtime}
	} else {
		return nil
	}
}

// Type of virtual network interface card (private, public)
func (n nic) Type() string {
	if n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return NIC_private
	}
	if n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return NIC_public
	}
	return ""
}
