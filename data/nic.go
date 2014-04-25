// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

// IPv4 describes properties of IPv4 address
type IPv4 struct {
	Conf string   `json:"conf"`
	IP   Resource `json:"ip"`
}

// RuntimeNetworkIO describes runtime I/O statistic for network interface card at runtime
type RuntimeNetworkIO struct {
	BytesRecv   int64 `json:"bytes_recv"`
	BytesSent   int64 `json:"bytes_sent"`
	PacketsRecv int64 `json:"packets_recv"`
	PacketsSent int64 `json:"packets_sent"`
}

// RuntimeNetwork describes properties of network interface card at runtime
type RuntimeNetwork struct {
	InterfaceType string           `json:"interface_type"`
	IO            RuntimeNetworkIO `json:"io"`
	IPv4          *Resource        `json:"ip_v4"`
}

// NIC describes properties of network interface card
type NIC struct {
	IPv4    *IPv4           `json:"ip_v4_conf"`
	Model   string          `json:"model"`
	MAC     string          `json:"mac"`
	VLAN    *Resource       `json:"vlan"`
	Runtime *RuntimeNetwork `json:"runtime"`
}
