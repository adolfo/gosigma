// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "github.com/Altoros/gosigma/data"

// A ServerDrive represents drive, connected to server instance
type ServerDrive struct {
	client *Client
	obj    *data.ServerDrive
}

// BootOrder of drive
func (sd ServerDrive) BootOrder() int {
	return sd.obj.BootOrder
}

// Channel of drive
func (sd ServerDrive) Channel() string {
	return sd.obj.Channel
}

// Device name of drive
func (sd ServerDrive) Device() string {
	return sd.obj.Device
}

// Drive object. Note, returned Drive object carries only UUID and URI, so it needs
// to perform Drive.Refresh to access additional information.
func (sd ServerDrive) Drive() Drive {
	obj := data.Drive{DriveRecord: data.DriveRecord{Resource: sd.obj.Drive}}
	return Drive{sd.client, &obj}
}

/*
// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (sd ServerDrive) String() string {
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
