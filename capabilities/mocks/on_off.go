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

func (m *OnOff) Status(c context.Context) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}

func (m *OnOff) On(c context.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *OnOff) Off(c context.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *OnOff) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *OnOff) Name() string {
	args := m.Called()
	return args.String(0)
}
