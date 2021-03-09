package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.EnumerateDevice = (*EnumerateDevice)(nil)

type EnumerateDevice struct {
	mock.Mock
}

func (m *EnumerateDevice) Status(c context.Context, d da.Device) (capabilities.EnumerationStatus, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.EnumerationStatus), args.Error(1)
}

func (m *EnumerateDevice) Enumerate(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}

func (m *EnumerateDevice) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *EnumerateDevice) Name() string {
	args := m.Called()
	return args.String(0)
}
