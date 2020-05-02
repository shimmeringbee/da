package da

type DeviceAdded struct {
	Device
}

type DeviceRemoved struct {
	Device
}

type DeviceStateChange struct {
	Device
}
