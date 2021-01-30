package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// WithLastUpdateTime allows a capability to optionally support providing the last time it received an update from
// the device with regards to the capability, it does not mean the values changed at that time. See LastChangeTime.
type WithLastUpdateTime interface {
	// LastUpdateTime returns the last time the capability received an update from the device. If the device supports
	// this feature, but has yet to receive an update, then time.isZero will be true.
	LastUpdateTime(context.Context, da.Device) (time.Time, error)
}

// WithLastChangeTime allows a capability to optionally support providing the last time the capability state meaningfully
// changed. For example if two updates for temperature match, then the second would not trigger a change of this value.
type WithLastChangeTime interface {
	// LastChangeTime returns the last time the capability received an update that caused the state to meaningfully change
	// from the device. If the device supports this feature, but has yet to receive an update, then time.isZero will be
	// true.
	LastChangeTime(context.Context, da.Device) (time.Time, error)
}
