// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"
	"time"

	"github.com/Altoros/gosigma/data"
)

const (
	// ServerStopped defines constant for stopped instance state
	ServerStopped = "stopped"
	// ServerStarting defines constant for starting instance state
	ServerStarting = "starting"
	// ServerRunning defines constant for running instance state
	ServerRunning = "running"
	// ServerStopping defines constant for stopping instance state
	ServerStopping = "stopping"
	// ServerUnavailable defines constant for unavailable instance state
	ServerUnavailable = "unavailable"
)

const (
	// RecurseNothing defines constant to remove server and leave all attached disks and CDROMs.
	RecurseNothing = ""
	// RecurseAllDrives defines constant to remove server and all attached drives regardless of media type they have.
	RecurseAllDrives = "all_drives"
	// RecurseDisks defines constant to remove server and all attached drives having media type "disk".
	RecurseDisks = "disks"
	// RecurseCDROMs defines constant to remove server and all attached drives having media type "cdrom".
	RecurseCDROMs = "cdroms"
)

// A Server represents server instance in CloudSigma account
type Server struct {
	client *Client
	obj    *data.Server
}

// Name of server instance
func (s Server) Name() string { return s.obj.Name }

// URI of server instance
func (s Server) URI() string { return s.obj.URI }

// Status of server instance
func (s Server) Status() string { return s.obj.Status }

// UUID of server instance
func (s Server) UUID() string { return s.obj.UUID }

// Cpu frequency in MHz
func (s Server) Cpu() int64 { return s.obj.CPU }

// Mem capacity in bytes
func (s Server) Mem() int64 { return s.obj.Mem }

// Get meta-information value stored in the server instance
func (s Server) Get(key string) (v string, ok bool) {
	v, ok = s.obj.Meta[key]
	return
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (s Server) String() string {
	return fmt.Sprintf("{Name: %q\nURI: %q\nStatus: %s\nUUID: %q}",
		s.Name(), s.URI(), s.Status(), s.UUID())
}

// Refresh information about server instance
func (s *Server) Refresh() error {
	obj, err := s.client.getServer(s.UUID())
	if err != nil {
		return err
	}
	s.obj = obj
	return nil
}

// Start server instance. This method does not check current server status,
// start command is issued to the endpoint in case of any value cached in Status().
func (s Server) Start() error {
	return s.client.startServer(s.UUID(), nil)
}

// Stop server instance. This method does not check current server status,
// stop command is issued to the endpoint in case of any value cached in Status().
func (s Server) Stop() error {
	return s.client.stopServer(s.UUID())
}

// Start server instance and waits for status ServerRunning with timeout
func (s *Server) StartWait() error {
	if err := s.Start(); err != nil {
		return err
	}
	return s.waitStatus(ServerRunning)
}

// Stop server instance and waits for status ServerStopped with timeout
func (s *Server) StopWait() error {
	if err := s.Stop(); err != nil {
		return err
	}
	return s.waitStatus(ServerStopped)
}

// Remove server instance
func (s Server) Remove(recurse string) error {
	return s.client.removeServer(s.UUID(), recurse)
}

func (s *Server) waitStatus(status string) error {
	var stop = false

	timeout := s.client.GetOperationTimeout()
	if timeout > 0 {
		timer := time.AfterFunc(timeout, func() { stop = true })
		defer timer.Stop()
	}

	for s.Status() != status {
		if err := s.Refresh(); err != nil {
			return err
		}
		if stop {
			return ErrOperationTimeout
		}
	}

	return nil
}
