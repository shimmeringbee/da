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

func (m *PressureSensor) Reading(c context.Context, d da.Device) ([]capabilities.PressureReading, error) {
	args := m.Called(c, d)
	return args.Get(0).([]capabilities.PressureReading), args.Error(1)
}
