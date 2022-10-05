package da

import (
	"github.com/shimmeringbee/zigbee"
)

//Capability identifier.
type Capability uint16

//Identifier provides an interface for any piece of data that uniquely identifies a device within a system.
type Identifier interface {
	// String from identifier.
	String() string
}

const DeviceDoesNotBelongToGatewayError = zigbee.Error("device does not belong to the gateway")
const DeviceIsNotGatewaySelfDeviceError = zigbee.Error("capability can only be called on gateway self device")
const DeviceIsGatewaySelfDeviceError = zigbee.Error("capability can not be called on gateway self device")
const DeviceDoesNotHaveCapability = zigbee.Error("device does not have capability")

func DeviceDoesNotBelongToGateway(gateway Gateway, device Device) bool {
	return device.Gateway() != gateway
}

func DeviceIsNotGatewaySelf(gateway Gateway, device Device) bool {
	return gateway.Self().Identifier() != device.Identifier()
}

func DeviceCapabilityCheck(device Device, gateway Gateway, capability Capability) error {
	if device.Gateway() != gateway {
		return DeviceDoesNotBelongToGatewayError
	}

	if !device.HasCapability(capability) {
		return DeviceDoesNotHaveCapability
	}

	return nil
}

type BasicCapability interface {
	Capability() Capability
	Name() string
}
