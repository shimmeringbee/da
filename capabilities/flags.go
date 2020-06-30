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
	OnOffFlag = Capability(0x1000)

	/* Developer Functionality. */
	LocalDebugFlag          = Capability(0xfff0)
	RemoteDebugFlag         = Capability(0xfff1)
	MessageCaptureDebugFlag = Capability(0xfff2)
)
