package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.HasProductInformation = (*HasProductInformation)(nil)

type HasProductInformation struct {
	mock.Mock
}

func (m *HasProductInformation) ProductInformation(c context.Context, d da.Device) (capabilities.ProductInformation, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.ProductInformation), args.Error(1)
}
