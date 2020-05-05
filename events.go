package da

// Device has been added to the gateway.
type DeviceAdded struct {
	// Device added.
	Device
}

// Device has been removed from gateway.
type DeviceRemoved struct {
	// Device removed.
	Device
}

// Device state has changed.
type DeviceStateChange struct {
	// Device that state has change don.
	Device
}
