package mocks

import (
	"context"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.Identify = (*Identify)(nil)

type Identify struct {
	mock.Mock
}

func (m *Identify) Identify(c context.Context, duration time.Duration) error {
	return m.Called(c, duration).Error(0)
}

func (m *Identify) Status(c context.Context) (capabilities.IdentifyState, error) {
	args := m.Called(c)
	return args.Get(0).(capabilities.IdentifyState), args.Error(1)
}
