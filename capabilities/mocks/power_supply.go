package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.PowerSupply = (*PowerSupply)(nil)

type PowerSupply struct {
	mock.Mock
}

func (m *PowerSupply) Status(c context.Context, d da.Device) (capabilities.PowerStatus, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.PowerStatus), args.Error(1)
}
