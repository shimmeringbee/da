package da

// DeviceAdded has been added to the gateway. Listen for CapabilityAdded to know what abilities are added to the
// Device.
type DeviceAdded struct {
	// Device added.
	Device Device
}

// DeviceRemoved has been removed from gateway.
type DeviceRemoved struct {
	// Device removed.
	Device Device
}

// CapabilityAdded signals a new capability has been added to the device.
type CapabilityAdded struct {
	// Device with new capability.
	Device Device
	// Capability being added.
	Capability Capability
}

// CapabilityRemoved signals a capability is being removed from the device.
type CapabilityRemoved struct {
	// Device with new capability.
	Device Device
	// Capability being added.
	Capability Capability
}
