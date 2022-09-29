package da

// Device is a set of interfaces that define the identity and capabilities of an abstracted device.
type Device interface {
	// Gateway returns the parent gateway.
	Gateway() Gateway
	// Identifier returns a unique identifier for this device.
	Identifier() Identifier
	// Capabilities returns a slice of capabilities this Device supports.
	Capabilities() []Capability
	// HasCapability returns true if a device supports the provided capability.
	HasCapability(Capability) bool
}

// BaseDevice is a concrete implementation of a device that is passed to consumers of da, implementers of a da may
// build upon this or just implement the Device interface above.
//
// Deprecated: BaseDevice should not be used, a da implementation specific version should be created.
type BaseDevice struct {
	// Gateway of the device.
	DeviceGateway Gateway
	// Identifier of device.
	DeviceIdentifier Identifier
	// Capabilities of device.
	DeviceCapabilities []Capability
}

// Gateway returns the capability of the base device.
func (d BaseDevice) Gateway() Gateway {
	return d.DeviceGateway
}

// Identifier returns the identifier of the base device.
func (d BaseDevice) Identifier() Identifier {
	return d.DeviceIdentifier
}

// Capabilities returns a slice of all capabilities this device has.
func (d BaseDevice) Capabilities() []Capability {
	return d.DeviceCapabilities
}

// HasCapability returns true if the capability provided is present in the base devices
// capability list.
func (d BaseDevice) HasCapability(c Capability) bool {
	for _, dC := range d.DeviceCapabilities {
		if dC == c {
			return true
		}
	}

	return false
}
