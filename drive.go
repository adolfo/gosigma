// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/Altoros/gosigma/data"
)

const (
	// DriveUnmounted defines constant for unmounted drive status
	DriveUnmounted = "unmounted"
	// DriveCreating defines constant for creating drive status
	DriveCreating = "creating"
	// DriveResizing defines constant for resizing drive status
	DriveResizing = "resizing"
	// DriveCloningDst defines constant for drive cloning status
	DriveCloningDst = "cloning_dst"
	// ... may be another values here, contact CloudSigma devs
)

const (
	// MediaCdrom defines media type for cdrom drives
	MediaCdrom = "cdrom"
	// MediaDisk defines media type for disk drives
	MediaDisk = "disk"
)

// A Drive represents drive instance in CloudSigma account
type Drive struct {
	client *Client
	obj    *data.Drive
}

// Name of drive instance
func (d Drive) Name() string { return d.obj.Name }

// URI of drive instance
func (d Drive) URI() string { return d.obj.URI }

// Status of drive instance
func (d Drive) Status() string { return d.obj.Status }

// UUID of drive instance
func (d Drive) UUID() string { return d.obj.UUID }

// Media of drive instance
func (d Drive) Media() string { return d.obj.Media }

// StorageType of drive instance
func (d Drive) StorageType() string { return d.obj.StorageType }

// Size of drive in bytes
func (d Drive) Size() int64 { return d.obj.Size }

// Get meta-information value stored in the drive instance
func (d Drive) Get(key string) (v string, ok bool) {
	v, ok = d.obj.Meta[key]
	return
}

// Jobs for this drive instance.
// Every job object in resulting slice carries only UUID and URI.
// To obtain additional information for job, one should use Job.Refresh() method
// to query cloud for detailed job information.
func (d Drive) Jobs() []Job {
	r := make([]Job, 0, len(d.obj.Jobs))
	for _, j := range d.obj.Jobs {
		job := Job{d.client, &data.Job{Resource: j}}
		r = append(r, job)
	}
	return r
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (d Drive) String() string {
	return fmt.Sprintf("{Name: %q\nURI: %q\nStatus: %s\nUUID: %q\nSize: %d\nMedia: %s\nStorage: %s}",
		d.Name(), d.URI(), d.Status(), d.UUID(), d.Size(), d.Media(), d.StorageType())
}

// Refresh information about drive instance
func (d *Drive) Refresh() error {
	obj, err := d.client.getDrive(d.UUID())
	if err != nil {
		return err
	}
	d.obj = obj
	return nil
}

// CloneParams defines attributes for drive cloning operation
type CloneParams struct {
	Affinities []string
	Media      string
	Name       string
}

func (c *CloneParams) makeJsonReader() (io.Reader, error) {
	if c == nil {
		return nil, nil
	}

	var m = make(map[string]interface{})
	if len(c.Affinities) > 0 {
		m["affinities"] = c.Affinities
	}
	if c.Media != "" {
		m["media"] = c.Media
	}
	if c.Name != "" {
		m["name"] = c.Name
	}

	bb, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bb), nil
}

// Clone drive instance.
func (d Drive) Clone(params *CloneParams, avoid []string) (Drive, error) {
	objs, err := d.client.cloneDrive(d.UUID(), params, avoid)

	if err != nil {
		return Drive{}, err
	}

	if len(objs) == 0 {
		return Drive{}, errors.New("No object was returned from server")
	}

	drv := Drive{d.client, &objs[0]}

	return drv, nil
}
