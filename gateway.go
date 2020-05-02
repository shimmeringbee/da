package da

type Gateway interface {
	ReadEvent() (interface{}, error)

	Capability(Capability) interface{}

	Self() Device
	Devices() []Device
}
