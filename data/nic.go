// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

// IPv4 describes properties of IPv4 address
type IPv4 struct {
	Conf string   `json:"conf"`
	IP   Resource `json:"ip"`
}

// NIC describes properties of network interface card
type NIC struct {
	IPv4  IPv4     `json:"ip_v4_conf"`
	Model string   `json:"model"`
	MAC   string   `json:"mac"`
	VLAN  Resource `json:"vlan"`
}
