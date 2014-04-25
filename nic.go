// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

const (
	NIC_vlan = "vlan"
	NIC_ip   = "ip"
)

// A NIC represents network interface card instance in CloudSigma server instance
type NIC struct {
	client *Client
	obj    *data.NIC
}

// Type of virtual network interface card (vlan or ip)
func (n NIC) Type() string {
	if n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return "vlan"
	}
	if n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return "ip"
	}
	return ""
}

// Conf returns type of network interface card configuration
func (n NIC) Conf() string {
	if n.obj.VLAN != nil && n.obj.VLAN.UUID != "" {
		return "vlan"
	}
	if n.obj.IPv4 != nil && n.obj.IPv4.Conf != "" {
		return n.obj.IPv4.Conf
	}
	return ""
}

// Model of virtual network interface card
func (n NIC) Model() string {
	return n.obj.Model
}

// MAC address
func (n NIC) MAC() string {
	return n.obj.MAC
}

// Runtime returns runtime information for network interface card or nil if stopped
func (n NIC) Runtime() *RuntimeNIC {
	if n.obj == nil || n.obj.Runtime == nil {
		return nil
	}
	return &RuntimeNIC{*n.obj.Runtime}
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (n NIC) String() string {
	return fmt.Sprintf(`{Type: %q, Conf: %q, Model: %s, MAC: %q}`,
		n.Type(), n.Conf(), n.Model(), n.MAC())
}
