// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"testing"

	"github.com/Altoros/gosigma/data"
)

func TestNIC_Empty(t *testing.T) {
	var n NIC
	if v := n.Type(); v != "" {
		t.Errorf("invalid type %q, must be empty", v)
	}
	if v := n.Conf(); v != "" {
		t.Errorf("invalid conf %q, must be empty", v)
	}
	if v := n.Model(); v != "" {
		t.Errorf("invalid model %q, must be empty", v)
	}
	if v := n.MAC(); v != "" {
		t.Errorf("invalid MAC %q, must be empty", v)
	}
	if v := n.Address(); v != "" {
		t.Errorf("invalid address %q, must be empty", v)
	}

	const str = `{Type: "", Conf: "", Model: "", MAC: "", Address: ""}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}

	r := n.Runtime()
	if r.obj != nil {
		t.Error("invalid runtime object")
	}
}

func TestNIC_DataIP(t *testing.T) {
	var d = data.NIC{
		IPv4: &data.IPv4{
			Conf: "static",
			IP:   &data.Resource{"/api/2.0/ips/31.171.246.37/", "31.171.246.37"},
		},
		Model: "virtio",
		MAC:   "22:40:85:4f:d3:ce",
	}
	var n = NIC{obj: &d}

	if v := n.Type(); v != NIC_public {
		t.Errorf("invalid type %q, must be %s", v, NIC_public)
	}
	if v := n.Conf(); v != "static" {
		t.Errorf("invalid conf %q, must be static", v)
	}
	if v := n.Model(); v != "virtio" {
		t.Errorf("invalid model %q, must be virtio", v)
	}
	if v := n.MAC(); v != "22:40:85:4f:d3:ce" {
		t.Errorf("invalid MAC %q, must be 22:40:85:4f:d3:ce", v)
	}
	if v := n.Address(); v != "31.171.246.37" {
		t.Errorf("invalid address %q, must be 31.171.246.37", v)
	}

	const str = `{Type: "public", Conf: "static", Model: "virtio", MAC: "22:40:85:4f:d3:ce", Address: "31.171.246.37"}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}

	r := n.Runtime()
	if r.obj != nil {
		t.Error("invalid runtime object")
	}
}

func TestNIC_DataVLan(t *testing.T) {
	var d = data.NIC{
		Model: "virtio",
		MAC:   "22:40:85:4f:d3:ce",
		VLAN: &data.Resource{
			"/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
			"5bc05e7e-6555-4f40-add8-3b8e91447702",
		},
	}
	var n = NIC{obj: &d}

	if v := n.Type(); v != NIC_private {
		t.Errorf("invalid type %q, must be %s", v, NIC_private)
	}
	if v := n.Conf(); v != NIC_private {
		t.Errorf("invalid conf %q, must be %s", v, NIC_private)
	}
	if v := n.Model(); v != "virtio" {
		t.Errorf("invalid model %q, must be virtio", v)
	}
	if v := n.MAC(); v != "22:40:85:4f:d3:ce" {
		t.Errorf("invalid MAC %q, must be 22:40:85:4f:d3:ce", v)
	}
	if v := n.Address(); v != "5bc05e7e-6555-4f40-add8-3b8e91447702" {
		t.Errorf("invalid address %q, must be 5bc05e7e-6555-4f40-add8-3b8e91447702", v)
	}

	const str = `{Type: "private", Conf: "private", Model: "virtio", MAC: "22:40:85:4f:d3:ce", Address: "5bc05e7e-6555-4f40-add8-3b8e91447702"}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}

	r := n.Runtime()
	if r.obj != nil {
		t.Error("invalid runtime object")
	}
}
