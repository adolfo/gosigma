// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"
	"time"

	"github.com/Altoros/gosigma/data"
)

const (
	// JobStateStarted defines constant for started job state
	JobStateStarted = "started"
	// JobStateSuccess defines constant for success job state
	JobStateSuccess = "success"
)

// A Job represents job instance in CloudSigma account
type Job struct {
	client *Client
	obj    *data.Job
}

// URI of job instance
func (j Job) URI() string { return j.obj.URI }

// UUID of job instance
func (j Job) UUID() string { return j.obj.UUID }

// Children of this job instance
func (j Job) Children() []string {
	r := make([]string, len(j.obj.Children))
	copy(r, j.obj.Children)
	return r
}

// Created time of this job instance
func (j Job) Created() time.Time { return j.obj.Created }

// Progress of this job instance
func (j Job) Progress() int { return j.obj.Data.Progress }

// LastModified time of this job instance
func (j Job) LastModified() time.Time { return j.obj.LastModified }

// Operation of this job instance
func (j Job) Operation() string { return j.obj.Operation }

// Resources of this job instance
func (j Job) Resources() []string {
	r := make([]string, len(j.obj.Resources))
	copy(r, j.obj.Resources)
	return r
}

// State of this job instance
func (j Job) State() string { return j.obj.State }

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (j Job) String() string {
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

// Refresh information about job instance
func (j *Job) Refresh() error {
	obj, err := j.client.getJob(j.UUID())
	if err != nil {
		return err
	}
	j.obj = obj
	return nil
}
