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

// A NIC represents network interface card instance in CloudSigma server instance
type NIC struct {
	client *Client
	obj    *data.NIC
}

// Type of virtual network interface card (private, public)
func (n NIC) Type() string {
	if n.obj != nil && n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return NIC_private
	}
	if n.obj != nil && n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return NIC_public
	}
	return ""
}

// Conf returns type of network interface card configuration. 'private' for NIC_vlan type,
// 'static', 'dhcp', 'manual' for NIC_ip
func (n NIC) Conf() string {
	if n.obj != nil && n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return NIC_private
	}
	if n.obj != nil && n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return n.obj.IPv4.Conf
	}
	return ""
}

// Address returns address configured for network interface card
func (n NIC) Address() string {
	if n.obj != nil && n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return n.obj.VLAN.UUID
	}
	if n.obj != nil && n.obj.IPv4 != nil && n.obj.IPv4.IP.UUID != "" {
		return n.obj.IPv4.IP.UUID
	}
	return ""
}

// Model of virtual network interface card
func (n NIC) Model() string {
	if n.obj != nil {
		return n.obj.Model
	} else {
		return ""
	}
}

// MAC address
func (n NIC) MAC() string {
	if n.obj != nil {
		return n.obj.MAC
	} else {
		return ""
	}
}

// Runtime returns runtime information for network interface card or nil if stopped
func (n NIC) Runtime() RuntimeNIC {
	if n.obj != nil {
		return RuntimeNIC{n.obj.Runtime}
	} else {
		return RuntimeNIC{}
	}
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (n NIC) String() string {
	return fmt.Sprintf(`{Type: %q, Conf: %q, Model: %q, MAC: %q, Address: %q}`,
		n.Type(), n.Conf(), n.Model(), n.MAC(), n.Address())
}
