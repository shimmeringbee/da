package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// LocalDebug is a capability which provides debug information about devices, this includes information that is already
// present in the device abstraction, returned information will always be in a format specific to the abstraction which
// may be marshalled as JSON.
type LocalDebug interface {
	// Start returns all locally kept information about the device.
	Start(context.Context) error
}

// LocalDebugStart sent to inform consumers that local debug is beginning on a device.
type LocalDebugStart struct {
	// Device that local debug started on.
	Device da.Device
}

// LocalDebugSuccess sent to inform consumers that local debug has completed successfully.
type LocalDebugSuccess struct {
	// Device that local debug succeeded on.
	Device da.Device
	// Media type of debug information, assuming JSON marshal.
	MediaType string
	// Debug payload.
	Debug interface{}
}

// LocalDebugFailure sent to inform consumers that local debug has failed.
type LocalDebugFailure struct {
	// Device that local debug failed on.
	Device da.Device
	// Error describing failure to produce local debug
	Error error
}
