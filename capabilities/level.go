package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"time"
)

// Level is a capability that provides the ability for a device to set the level of an output, this
// could be brightness, fan intensity, blind height, desk height, etc.
type Level interface {
	// ChangeTo changes the devices current output level to the new required value, over a set duration.
	ChangeTo(context.Context, Device, float64, time.Duration) error
}
