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
	// StartMessageCapture begins capturing messages, it will buffer `number` messages, once that many messages are
	// captured, if `wrap` is true then the earliest messages will be replaced, if false capturing will end. This
	// function also clears the buffer at the start.
	StartMessageCapture(context context.Context, device Device) error

	// StopMessageCapture stops capturing messages, it does not wipe the buffer.
	StopMessageCapture(contact context.Context, device Device) error
}

// Event sent to inform consumers that message capture is beginning on a device.
type MessageCaptureStart struct {
	// Device that message capture started on.
	Device
}

// Event sent to inform consumers that message capture is ended on a device.
type MessageCapture struct {
	// Device that message capture occurred on.
	Device
	// Captured Message.
	CapturedMessage
}

// Event sent to inform consumers that message capture is ended on a device.
type MessageCaptureStop struct {
	// Device that message capture stopped on.
	Device
}
