package da

// Device has been added to the gateway, it is unlikely to be ready for usage at this stage. A DeviceLoaded or
// EnumerateDeviceSuccess will follow to mark load/enumeration completion.
type DeviceAdded struct {
	// Device added.
	Device
}

// Device has been loaded into the gateway. The intent of this event is to inform the da consumer that a device has been
// loaded from configuration or a previous execution. da consumers may wish to treat this similar to
// EnumerateDeviceSuccess.
type DeviceLoaded struct {
	// Device loaded.
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
