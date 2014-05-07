// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import "github.com/Altoros/gosigma/data"

func (d drive) clone(params CloneParams, avoid []string) (*data.Drive, error) {
	obj, err := d.client.cloneDrive(d.UUID(), d.Library(), params, avoid)
	if err != nil {
		return nil, err
	}

	if d.Library() == LibraryMedia {
		obj.LibraryDrive = d.obj.LibraryDrive
	}

	return obj, nil
}
