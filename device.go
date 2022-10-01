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

// SimpleDevice is a concrete implementation of a device that is passed to consumers of da, implementers of a da may
// build upon this or just implement the Device interface above.
type SimpleDevice struct {
	// Gateway of the device.
	DeviceGateway Gateway
	// Identifier of device.
	DeviceIdentifier Identifier
	// Capabilities of device.
	DeviceCapabilities []Capability
}

// Gateway returns the capability of the base device.
func (s SimpleDevice) Gateway() Gateway {
	return s.DeviceGateway
}

// Identifier returns the identifier of the base device.
func (s SimpleDevice) Identifier() Identifier {
	return s.DeviceIdentifier
}

// Capabilities returns a slice of all capabilities this device has.
func (s SimpleDevice) Capabilities() []Capability {
	return s.DeviceCapabilities
}

// HasCapability returns true if the capability provided is present in the devices capabilities list.
func (s SimpleDevice) HasCapability(needle Capability) bool {
	for _, straw := range s.DeviceCapabilities {
		if straw == needle {
			return true
		}
	}

	return false
}
