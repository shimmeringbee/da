package da

// Capability identifier.
type Capability uint16

// Identifier provides an interface for any piece of data that uniquely identifies a device within a system.
type Identifier interface {
	// String from identifier.
	String() string
}

type BasicCapability interface {
	Capability() Capability
	Name() string
}
