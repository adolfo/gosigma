// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"testing"
	"time"

	"github.com/Altoros/gosigma/data"
)

func TestJobString(t *testing.T) {
	j := &job{obj: &data.Job{}}
	if s := j.String(); s != `{UUID: "", Operation: , State: , Progress: 0, Resources: []}` {
		t.Errorf("invalid Job.String(): `%s`", s)
	}

	j.obj.UUID = "uuid"
	if s := j.String(); s != `{UUID: "uuid", Operation: , State: , Progress: 0, Resources: []}` {
		t.Errorf("invalid Job.String(): `%s`", s)
	}

	j.obj.Operation = "operation"
	if s := j.String(); s != `{UUID: "uuid", Operation: operation, State: , Progress: 0, Resources: []}` {
		t.Errorf("invalid Job.String(): `%s`", s)
	}

	j.obj.State = JobStateStarted
	if s := j.String(); s != `{UUID: "uuid", Operation: operation, State: started, Progress: 0, Resources: []}` {
		t.Errorf("invalid Job.String(): `%s`", s)
	}

	j.obj.Data.Progress = 99
	if s := j.String(); s != `{UUID: "uuid", Operation: operation, State: started, Progress: 99, Resources: []}` {
		t.Errorf("invalid Job.String(): `%s`", s)
	}
}

func TestJobChildren(t *testing.T) {
	j := &job{obj: &data.Job{}}
	if c := j.Children(); c == nil || len(c) != 0 {
		t.Errorf("invalid Job.Children(): %v", c)
	}

	j.obj.Children = append(j.obj.Children, "child-0")
	if c := j.Children(); c == nil || len(c) != 1 || c[0] != "child-0" {
		t.Errorf("invalid Job.Children(): %v", c)
	}

	j.obj.Children = append(j.obj.Children, "child-1")
	if c := j.Children(); c == nil || len(c) != 2 || c[0] != "child-0" || c[1] != "child-1" {
		t.Errorf("invalid Job.Children(): %v", c)
	}
}

func TestJobCreated(t *testing.T) {
	j := &job{obj: &data.Job{}}
	if c := j.Created(); c != (time.Time{}) {
		t.Errorf("invalid Job.Time(): %v", c)
	}

	j.obj.Created = time.Unix(100, 200)
	if c := j.Created(); c != time.Unix(100, 200) {
		t.Errorf("invalid Job.Time(): %v", c)
	}
}

func TestJobLastModified(t *testing.T) {
	j := &job{obj: &data.Job{}}
	if c := j.LastModified(); c != (time.Time{}) {
		t.Errorf("invalid Job.LastModified(): %v", c)
	}

	j.obj.LastModified = time.Unix(100, 200)
	if c := j.LastModified(); c != time.Unix(100, 200) {
		t.Errorf("invalid Job.LastModified(): %v", c)
	}
}

func TestJobRefresh(t *testing.T) {
	//mock.ResetJobs()
}
