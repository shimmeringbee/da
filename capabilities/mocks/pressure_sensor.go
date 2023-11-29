package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.PressureSensor = (*PressureSensor)(nil)

type PressureSensor struct {
	mock.Mock
}

func (m *PressureSensor) Reading(c context.Context) ([]capabilities.PressureReading, error) {
	args := m.Called(c)
	return args.Get(0).([]capabilities.PressureReading), args.Error(1)
}

func (m *PressureSensor) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *PressureSensor) Name() string {
	args := m.Called()
	return args.String(0)
}
