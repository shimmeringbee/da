package da

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDevice_HasCapability(t *testing.T) {
	capability := Capability(0x01)

	t.Run("it returns true if a capability exists", func(t *testing.T) {
		device := Device{
			Capabilities: []Capability{capability},
		}

		assert.True(t, device.HasCapability(capability))
	})

	t.Run("it returns false if a capability exists", func(t *testing.T) {
		device := Device{
			Capabilities: []Capability{},
		}

		assert.False(t, device.HasCapability(capability))
	})
}
