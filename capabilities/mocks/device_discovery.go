package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.DeviceDiscovery = (*DeviceDiscovery)(nil)

type DeviceDiscovery struct {
	mock.Mock
}

func (m *DeviceDiscovery) Status(c context.Context, d da.Device) (capabilities.DeviceDiscoveryStatus, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.DeviceDiscoveryStatus), args.Error(1)
}

func (m *DeviceDiscovery) Enable(c context.Context, d da.Device, du time.Duration) error {
	args := m.Called(c, d, du)
	return args.Error(0)
}

func (m *DeviceDiscovery) Disable(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}

func (m *DeviceDiscovery) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *DeviceDiscovery) Name() string {
	args := m.Called()
	return args.String(0)
}
