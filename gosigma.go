// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"errors"
	"fmt"
)

type VersionNum struct {
	Major int
	Minor int
	Micro int
}

func (v VersionNum) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Micro)
}

func VersionNumber() VersionNum {
	return VersionNum{Major: 0, Minor: 1, Micro: 0}
}

func Version() string {
	return VersionNumber().String()
}

const DefaultRegion string = "zrh"

func GetRegionEndpoint(r string) (string, error) {
	switch r {
	case "zrh":
		return "https://zrh.cloudsigma.com/api/2.0/", nil
	case "lvs":
		return "https://lvs.cloudsigma.com/api/2.0/", nil
	default:
		return "", errors.New("unknown CloudSigma region")
	}
}
