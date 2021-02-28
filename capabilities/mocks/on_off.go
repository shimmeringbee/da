package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.OnOff = (*OnOff)(nil)

type OnOff struct {
	mock.Mock
}

func (m *OnOff) Status(c context.Context, d da.Device) (bool, error) {
	args := m.Called(c, d)
	return args.Bool(0), args.Error(1)
}

func (m *OnOff) On(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}

func (m *OnOff) Off(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}
