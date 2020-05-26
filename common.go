package da

import "errors"

// Capability identifier.
type Capability uint16

// Interface for a device identifier.
type Identifier interface {
	// String from identifier.
	String() string
}

var DeviceDoesNotBelongToGatewayError = errors.New("device does not belong to the gateway")
var DeviceIsNotGatewaySelfDeviceError = errors.New("capability can only be called on gateway self device")
var DeviceIsGatewaySelfDeviceError = errors.New("capability can not be called on gateway self device")
var DeviceDoesNotHaveCapability = errors.New("device does not have capability")

func DeviceDoesNotBelongToGateway(gateway Gateway, device Device) bool {
	return device.Gateway != gateway
}

func DeviceIsNotGatewaySelf(gateway Gateway, device Device) bool {
	return gateway.Self().Identifier != device.Identifier
}
