// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

// Meta describes properties of dataset
type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

// Resource describes properties of linked resource
type Resource struct {
	URI  string `json:"resource_uri"`
	UUID string `json:"uuid"`
}
