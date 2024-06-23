package mocks

import (
	"context"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.DeviceWorkarounds = (*DeviceWorkarounds)(nil)

type DeviceWorkarounds struct {
	mock.Mock
}

func (m *DeviceWorkarounds) Enabled(c context.Context) ([]string, error) {
	args := m.Called(c)
	return args.Get(0).([]string), args.Error(1)
}
