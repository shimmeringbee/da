package capabilities

import "context"

// DeviceWorkarounds is a capability that exposes implementation details from the da. Specifically any
// modes it is having to operate in.
type DeviceWorkarounds interface {
	// Enabled retrieves a list of the workarounds modes that are active for the device.
	Enabled(context.Context) ([]string, error)
}
