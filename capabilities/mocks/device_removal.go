package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.DeviceRemoval = (*DeviceRemoval)(nil)

type DeviceRemoval struct {
	mock.Mock
}

func (m *DeviceRemoval) Remove(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}
