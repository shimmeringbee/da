package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"time"
)

// CapturedMessage represents a message sent from the device abstraction to a device.
type CapturedMessage struct {
	// Time message was captured.
	Timestamp time.Time
	// Source identifier of message.
	Source Identifier
	// Destination identifier of message.
	Destination Identifier
	// Outbound messages (to devices).
	Outbound bool
	// JSON marshallable copy of payload.
	Payload interface{}
}

// MessageCaptureDebug is a capability which allows capture of the internal protocol of the device abstraction, returned
// information will always be in a format specific to the abstraction which may be marshalled as JSON.
type MessageCaptureDebug interface {
	// StartMessageCapture begins capturing messages for a device, messages continue to be sent until StopMessageCapture
	// or the device is no longer present.
	StartMessageCapture(context context.Context, device Device) error

	// StopMessageCapture stops capturing messages.
	StopMessageCapture(context context.Context, device Device) error

	// EnableMessageCaptureOnJoin enables automatic capturing of messages of new devices that join, only callable on
	// the gateways self device.
	EnableMessageCaptureOnJoin(context context.Context) error

	// DisableMessageCaptureOnJoin disables automatic capturing of messages of new devices that join, only callable on
	// the gateways self device.
	DisableMessageCaptureOnJoin(context context.Context) error
}

// Event sent to inform consumers that message capture is beginning on a device.
type MessageCaptureStart struct {
	// Device that message capture started on.
	Device Device
}

// Event sent to inform consumers that message capture is ended on a device.
type MessageCapture struct {
	// Device that message capture occurred on.
	Device Device
	// Captured Message.
	Message CapturedMessage
}

// Event sent to inform consumers that message capture is ended on a device.
type MessageCaptureStop struct {
	// Device that message capture stopped on.
	Device Device
}
