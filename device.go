package da

// Abstraction of a device available on a gateway
type Device struct {
	// Gateway which owns this device.
	Gateway Gateway
	// Identifier which is used to refer to this device.
	Identifier Identifier
	// Capabilities this device supports.
	Capabilities []Capability
}

// Check to see if a device supports a particular capability.
func (d Device) HasCapability(capability Capability) bool {
	for _, haystack := range d.Capabilities {
		if haystack == capability {
			return true
		}
	}
	return false
}
