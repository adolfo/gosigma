// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "github.com/Altoros/gosigma/data"

type NetworkInterfaceType int

const (
	NetworkUnknown NetworkInterfaceType = iota
	NetworkDynamic
	NetworkStatic
	NetworkVLan
)

// A NIC represents network interface card instance in CloudSigma server instance
type NIC struct {
	client *Client
	obj    *data.NIC
}

// Type of network interface card
func (n NIC) Type() NetworkInterfaceType {
	switch n.obj.IPv4.Conf {
	case "dhcp":
		return NetworkDynamic
	case "static":
		return NetworkStatic
	}
	if n.obj.VLAN.UUID != "" {
		return NetworkVLan
	}
	return NetworkUnknown
}

// Model of virtual network interface card
func (n NIC) Model() string {
	return n.obj.Model
}

// MAC address
func (n NIC) MAC() string {
	return n.obj.MAC
}

// VLAN information
func (n NIC) VLAN() data.Resource {
	return n.obj.VLAN
}

// IPv4 information
func (n NIC) IPv4() data.IPv4 {
	return n.obj.IPv4
}

/*
// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (n NIC) String() string {
	return fmt.Sprintf(`{UUID: %q
Operation: %s
State: %s
Progress: %d,
Resources: %v}`,
		j.obj.UUID,
		j.obj.Operation,
		j.obj.State,
		j.obj.Data.Progress,
		j.obj.Resources)
}
*/
