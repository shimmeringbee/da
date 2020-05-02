package da

import "context"

type Gateway interface {
	ReadEvent(context.Context) (interface{}, error)

	Capability(Capability) interface{}

	Self() Device
	Devices() []Device

	Start() error
	Stop() error
}
