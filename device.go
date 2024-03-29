package da

// Device is a set of interfaces that define the identity and capabilities of an abstracted device.
type Device interface {
	// Gateway returns the parent gateway.
	Gateway() Gateway
	// Identifier returns a unique identifier for this device.
	Identifier() Identifier
	// Capabilities returns a slice of capabilities this Device supports.
	Capabilities() []Capability
	// Capability returns the concrete implementation of the capability.
	Capability(Capability) BasicCapability
}
