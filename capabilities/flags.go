package capabilities

import (
	. "github.com/shimmeringbee/da"
)

const (
	/* Basic capabilities to permit management of devices. */
	DeviceDiscoveryFlag       = Capability(0x0000)
	EnumerateDeviceFlag       = Capability(0x0001)
	HasProductInformationFlag = Capability(0x0002)

	/* Functional capabilities of devices. */
)
