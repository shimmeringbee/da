package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// CapturedMessage represents a message sent from the device abstraction to a device.
type CapturedMessage struct {
	// Time message was captured.
	Timestamp time.Time
	// Source identifier of message.
	Source da.Identifier
	// Destination identifier of message.
	Destination da.Identifier
	// Outbound messages (to devices).
	Outbound bool
	// JSON marshallable copy of payload.
	Payload interface{}
}

// MessageCaptureDebug is a capability which allows capture of the internal protocol of the device abstraction, returned
// information will always be in a format specific to the abstraction which may be marshalled as JSON.
type MessageCaptureDebug interface {
	// Start begins capturing messages for a device, messages continue to be sent until StopMessageCapture
	// or the device is no longer present.
	Start(context context.Context) error

	// Stop stops capturing messages.
	Stop(context context.Context) error

	// Status returns the state of device capturing.
	Status(context context.Context) (bool, error)

	// Get recently captured messages from device.
	Get(context context.Context) ([]MessageCapture, error)

	// Clear recently captured messages from device.
	Clear(context context.Context) error

	// EnableCaptureOnJoin enables automatic capturing of messages of new devices that join, only callable on
	// the gateways self device.
	EnableCaptureOnJoin(context context.Context) error

	// DisableCaptureOnJoin disables automatic capturing of messages of new devices that join, only callable on
	// the gateways self device.
	DisableCaptureOnJoin(context context.Context) error
}

// MessageCaptureStart sent to inform consumers that message capture is beginning on a device.
type MessageCaptureStart struct {
	// Device that message capture started on.
	Device da.Device
}

// MessageCapture sent to inform consumers that message capture is ended on a device.
type MessageCapture struct {
	// Device that message capture occurred on.
	Device da.Device
	// Captured Message.
	Message CapturedMessage
}

// MessageCaptureStop sent to inform consumers that message capture is ended on a device.
type MessageCaptureStop struct {
	// Device that message capture stopped on.
	Device da.Device
}
