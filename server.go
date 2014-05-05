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

// A Server interface represents server instance in CloudSigma account
type Server interface {
	// Convert to string
	fmt.Stringer

	// Context serial device enabled for server instance
	Context() bool

	// Cpu frequency in MHz
	Cpu() int64

	// Drives for this server instance
	Drives() []ServerDrive

	// Mem capacity in bytes
	Mem() int64

	// Name of server instance
	Name() string

	// NICs for this server instance
	NICs() []NIC

	// Status of server instance
	Status() string

	// URI of server instance
	URI() string

	// UUID of server instance
	UUID() string

	// VNCPassword to access the server
	VNCPassword() string

	// Get meta-information value stored in the server instance
	Get(key string) (string, bool)

	// Refresh information about server instance
	Refresh() error

	// Start server instance. This method does not check current server status,
	// start command is issued to the endpoint in case of any value cached in Status().
	Start() error

	// Stop server instance. This method does not check current server status,
	// stop command is issued to the endpoint in case of any value cached in Status().
	Stop() error

	// Start server instance and waits for status ServerRunning with timeout
	StartWait() error

	// Stop server instance and waits for status ServerStopped with timeout
	StopWait() error

	// Remove server instance
	Remove(recurse string) error
}

// A server implements server instance in CloudSigma account
type server struct {
	client *Client
	obj    *data.Server
}

var _ Server = (*server)(nil)

// String method implements fmt.Stringer interface
func (s server) String() string {
	return fmt.Sprintf("{Name: %q\nURI: %q\nStatus: %s\nUUID: %q}",
		s.Name(), s.URI(), s.Status(), s.UUID())
}

// Context serial device enabled for server instance
func (s server) Context() bool { return s.obj.Context }

// Cpu frequency in MHz
func (s server) Cpu() int64 { return s.obj.CPU }

// Drives for this server instance
func (s server) Drives() []ServerDrive {
	r := make([]ServerDrive, 0, len(s.obj.Drives))
	for i := range s.obj.Drives {
		drive := &serverDrive{s.client, &s.obj.Drives[i]}
		r = append(r, drive)
	}
	return r
}

// Mem capacity in bytes
func (s server) Mem() int64 { return s.obj.Mem }

// Name of server instance
func (s server) Name() string { return s.obj.Name }

// NICs for this server instance
func (s server) NICs() []NIC {
	r := make([]NIC, 0, len(s.obj.NICs))
	for i := range s.obj.NICs {
		n := nic{s.client, &s.obj.NICs[i]}
		r = append(r, n)
	}
	return r
}

// Status of server instance
func (s server) Status() string { return s.obj.Status }

// URI of server instance
func (s server) URI() string { return s.obj.URI }

// UUID of server instance
func (s server) UUID() string { return s.obj.UUID }

// VNCPassword to access the server
func (s server) VNCPassword() string { return s.obj.VNCPassword }

// Get meta-information value stored in the server instance
func (s server) Get(key string) (v string, ok bool) {
	v, ok = s.obj.Meta[key]
	return
}

// Refresh information about server instance
func (s *server) Refresh() error {
	obj, err := s.client.getServer(s.UUID())
	if err != nil {
		return err
	}
	s.obj = obj
	return nil
}

// Start server instance. This method does not check current server status,
// start command is issued to the endpoint in case of any value cached in Status().
func (s server) Start() error {
	return s.client.startServer(s.UUID(), nil)
}

// Stop server instance. This method does not check current server status,
// stop command is issued to the endpoint in case of any value cached in Status().
func (s server) Stop() error {
	return s.client.stopServer(s.UUID())
}

// Start server instance and waits for status ServerRunning with timeout
func (s *server) StartWait() error {
	if err := s.Start(); err != nil {
		return err
	}
	return s.waitStatus(ServerRunning)
}

// Stop server instance and waits for status ServerStopped with timeout
func (s *server) StopWait() error {
	if err := s.Stop(); err != nil {
		return err
	}
	return s.waitStatus(ServerStopped)
}

// Remove server instance
func (s server) Remove(recurse string) error {
	return s.client.removeServer(s.UUID(), recurse)
}

func (s *server) waitStatus(status string) error {
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
