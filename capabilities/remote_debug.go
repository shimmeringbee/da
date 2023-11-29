package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// RemoteDebug is a capability which provides debug information about devices, this includes information that is already
// present in the device abstraction as well as data returned from the device upon request. Returned information will
// always be in a format specific to the abstraction which may be marshalled as JSON.
type RemoteDebug interface {
	// Start returns all locally kept information about the device, and interogates the device for reasonable information
	// which may assist in the diagnosis of an issue.
	Start(context.Context) error
}

// RemoteDebugStart sent to inform consumers that remote debug is beginning on a device.
type RemoteDebugStart struct {
	// Device that remote debug started on.
	Device da.Device
}

// RemoteDebugSuccess sent to inform consumers that remote debug has completed successfully.
type RemoteDebugSuccess struct {
	// Device that remote debug succeeded on.
	Device da.Device
	// Media type of debug information, assuming JSON marshal.
	MediaType string
	// Debug payload.
	Debug interface{}
}

// RemoteDebugFailure sent to inform consumers that remote debug has failed.
type RemoteDebugFailure struct {
	// Device that remote debug failed on.
	Device da.Device
	// Error describing failure to produce local debug
	Error error
}
