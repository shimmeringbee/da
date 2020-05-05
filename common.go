package da

// Capability identifier.
type Capability uint16

// Interface for a device identifier.
type Identifier interface {
	// String from identifier.
	String() string
}
