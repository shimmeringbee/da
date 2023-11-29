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

func (m *DeviceRemoval) Remove(c context.Context, t capabilities.RemovalType) error {
	args := m.Called(c, t)
	return args.Error(0)
}

func (m *DeviceRemoval) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *DeviceRemoval) Name() string {
	args := m.Called()
	return args.String(0)
}
