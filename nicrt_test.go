// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"testing"

	"github.com/Altoros/gosigma/data"
)

func TestRuntimeNIC_Empty(t *testing.T) {
	var n RuntimeNIC
	if v := n.Type(); v != "" {
		t.Errorf("invalid type %q, must be empty", v)
	}
	if v := n.AddressIPv4(); v != "" {
		t.Errorf("invalid address %q, must be empty", v)
	}

	const str = `{Type: "", Address: ""}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}
}

func TestRuntimeNIC_Public(t *testing.T) {
	var obj = data.RuntimeNetwork{
		InterfaceType: "public",
		IPv4:          data.MakeIPResource("10.11.12.13"),
	}
	var n = RuntimeNIC{obj: &obj}
	if v := n.Type(); v != "public" {
		t.Errorf("invalid type %q, must be public", v)
	}
	if v := n.AddressIPv4(); v != "10.11.12.13" {
		t.Errorf("invalid address %q, must be empty", v)
	}

	const str = `{Type: "public", Address: "10.11.12.13"}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}
}

func TestRuntimeNIC_Private(t *testing.T) {
	var obj = data.RuntimeNetwork{
		InterfaceType: "private",
	}
	var n = RuntimeNIC{obj: &obj}
	if v := n.Type(); v != "private" {
		t.Errorf("invalid type %q, must be private", v)
	}
	if v := n.AddressIPv4(); v != "" {
		t.Errorf("invalid address %q, must be empty", v)
	}

	const str = `{Type: "private", Address: ""}`
	if v := n.String(); v != str {
		t.Errorf("invalid String() result: %q, must be %s", v, str)
	}
}
