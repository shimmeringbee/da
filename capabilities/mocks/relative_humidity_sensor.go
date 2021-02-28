package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.RelativeHumiditySensor = (*RelativeHumiditySensor)(nil)

type RelativeHumiditySensor struct {
	mock.Mock
}

func (m *RelativeHumiditySensor) Reading(c context.Context, d da.Device) ([]capabilities.RelativeHumidityReading, error) {
	args := m.Called(c, d)
	return args.Get(0).([]capabilities.RelativeHumidityReading), args.Error(1)
}
