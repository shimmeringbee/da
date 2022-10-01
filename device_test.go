package da

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleDevice_Parameters(t *testing.T) {
	t.Run("HasCapability returns true if capability is found", func(t *testing.T) {
		c := Capability(0x01)

		b := SimpleDevice{
			DeviceCapabilities: []Capability{c},
		}

		assert.True(t, b.HasCapability(c))
	})

	t.Run("HasCapability returns false if capability is not found", func(t *testing.T) {
		c := Capability(0x01)

		b := SimpleDevice{
			DeviceCapabilities: []Capability{},
		}

		assert.False(t, b.HasCapability(c))
	})
}
