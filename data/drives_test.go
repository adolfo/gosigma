// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDataDriversReaderFail(t *testing.T) {
	r := failReader{}

	if _, err := ReadDrive(r); err == nil || err.Error() != "test error" {
		t.Error("Fail")
	}

	if _, err := ReadDrives(r); err == nil || err.Error() != "test error" {
		t.Error("Fail")
	}
}

func TestDataDrivesUnmarshal(t *testing.T) {
	var dd Drives
	dd.Meta.Limit = 12345
	dd.Meta.Offset = 12345
	dd.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonDrivesData), &dd)
	if err != nil {
		t.Error(err)
	}

	verifyMeta(t, &dd.Meta, 0, 0, 9)

	for i := 0; i < len(drivesData); i++ {
		compareDrives(t, i, &dd.Objects[i], &drivesData[i])
	}
}

func TestDataDrivesDetailUnmarshal(t *testing.T) {
	var dd Drives
	dd.Meta.Limit = 12345
	dd.Meta.Offset = 12345
	dd.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonDrivesDetailData), &dd)
	if err != nil {
		t.Error(err)
	}

	verifyMeta(t, &dd.Meta, 0, 0, 9)

	for i := 0; i < len(drivesDetailData); i++ {
		compareDrives(t, i, &dd.Objects[i], &drivesDetailData[i])
	}
}

func TestDataDrivesReadDrives(t *testing.T) {
	dd, err := ReadDrives(strings.NewReader(jsonDrivesData))
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(drivesData); i++ {
		compareDrives(t, i, &dd[i], &drivesData[i])
	}
}

func TestDataDrivesReadDrive(t *testing.T) {
	d, err := ReadDrive(strings.NewReader(jsonDriveData))
	if err != nil {
		t.Error(err)
	}

	compareDrives(t, 0, d, &driveData)
}

func compareDrives(t *testing.T, i int, value, wants *Drive) {
	if value.DriveRecord != wants.DriveRecord {
		t.Errorf("DriveRecord error [%d]: found %#v, wants %#v", i, value.DriveRecord, wants.DriveRecord)
	}
	if len(value.Jobs) != len(wants.Jobs) {
		t.Errorf("Drive.Jobs error [%d]: found %#v, wants %#v", i, value.Jobs, wants.Jobs)
	}
	if value.Media != wants.Media {
		t.Errorf("Drive.Media error [%d]: found %#v, wants %#v", i, value.Media, wants.Media)
	}
	if len(value.Meta) != len(wants.Meta) {
		t.Errorf("Drive.Meta error [%d]: found %#v, wants %#v", i, value.Meta, wants.Meta)
	}
	if value.Size != wants.Size {
		t.Errorf("Drive.Size error [%d]: found %#v, wants %#v", i, value.Size, wants.Size)
	}
	if value.StorageType != wants.StorageType {
		t.Errorf("Drive.StorageType error [%d]: found %#v, wants %#v", i, value.StorageType, wants.StorageType)
	}
	if value.Name != wants.Name {
		t.Errorf("Drive.Name error [%d]: found %#v, wants %#v", i, value.Name, wants.Name)
	}
}
