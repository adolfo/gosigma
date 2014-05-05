// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"
	"strings"

	"github.com/Altoros/gosigma/data"
)

// A ServerDrive interface represents drive, connected to server instance
type ServerDrive interface {
	// Convert to string
	fmt.Stringer

	// BootOrder of drive
	BootOrder() int

	// Channel of drive
	Channel() string

	// Device name of drive
	Device() string

	// Drive object. Note, returned Drive object carries only UUID and URI, so it needs
	// to perform Drive.Refresh to access additional information.
	Drive() Drive

	// UUID of drive
	UUID() string
}

// A serverDrive implements drive, connected to server instance
type serverDrive struct {
	client *Client
	obj    *data.ServerDrive
}

var _ ServerDrive = serverDrive{}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (sd serverDrive) String() string {
	return fmt.Sprintf(`{BootOrder: %d, Channel: %q, Device: %q, UUID: %q}`,
		sd.BootOrder(), sd.Channel(), sd.Device(), sd.UUID())
}

// BootOrder of drive
func (sd serverDrive) BootOrder() int {
	if sd.obj != nil {
		return sd.obj.BootOrder
	} else {
		return 0
	}
}

// Channel of drive
func (sd serverDrive) Channel() string {
	if sd.obj != nil {
		return sd.obj.Channel
	} else {
		return ""
	}
}

// Device name of drive
func (sd serverDrive) Device() string {
	if sd.obj != nil {
		return sd.obj.Device
	} else {
		return ""
	}
}

// Drive object. Note, returned Drive object carries only UUID and URI, so it needs
// to perform Drive.Refresh to access additional information.
func (sd serverDrive) Drive() Drive {
	if sd.obj != nil {
		obj := data.Drive{Resource: sd.obj.Drive}
		libdrive := strings.Contains(sd.obj.Drive.UUID, "libdrives")
		if libdrive {
			return Drive{sd.client, &obj, LibraryMedia}
		} else {
			return Drive{sd.client, &obj, LibraryAccount}
		}
	} else {
		return Drive{sd.client, nil, false}
	}
}

// UUID of drive
func (sd serverDrive) UUID() string {
	if sd.obj != nil {
		return sd.obj.Drive.UUID
	} else {
		return ""
	}
}
