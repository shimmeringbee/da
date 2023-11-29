package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

type ProductInformationPresent uint

const (
	Manufacturer ProductInformationPresent = 1 << iota
	Name
	Serial
	Version
)

type ProductInformation struct {
	Present ProductInformationPresent

	Manufacturer string
	Name         string
	Serial       string
	Version      string
}

type HasProductInformation interface {
	ProductInformation(context.Context, da.Device) (ProductInformation, error)
}
