// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"
	"strings"

	"github.com/Altoros/gosigma/data"
)

// A ServerDrive represents drive, connected to server instance
type ServerDrive struct {
	client *Client
	obj    *data.ServerDrive
}

// BootOrder of drive
func (sd ServerDrive) BootOrder() int {
	if sd.obj != nil {
		return sd.obj.BootOrder
	} else {
		return 0
	}
}

// Channel of drive
func (sd ServerDrive) Channel() string {
	if sd.obj != nil {
		return sd.obj.Channel
	} else {
		return ""
	}
}

// Device name of drive
func (sd ServerDrive) Device() string {
	if sd.obj != nil {
		return sd.obj.Device
	} else {
		return ""
	}
}

// UUID of drive
func (sd ServerDrive) UUID() string {
	if sd.obj != nil {
		return sd.obj.Drive.UUID
	} else {
		return ""
	}
}

// Drive object. Note, returned Drive object carries only UUID and URI, so it needs
// to perform Drive.Refresh to access additional information.
func (sd ServerDrive) Drive() Drive {
	if sd.obj != nil {
		obj := data.Drive{DriveRecord: data.DriveRecord{Resource: sd.obj.Drive}}
		libdrive := strings.Contains(sd.obj.Drive.UUID, "libdrives")
		return Drive{sd.client, &obj, libdrive}
	} else {
		return Drive{sd.client, nil, false}
	}
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (sd ServerDrive) String() string {
	return fmt.Sprintf(`{BootOrder: %d, Channel: %q, Device: %q, UUID: %q}`,
		sd.BootOrder(), sd.Channel(), sd.Device(), sd.UUID())
}
