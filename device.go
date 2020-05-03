package da

type Device struct {
	Gateway      Gateway
	Identifier   Identifier
	Capabilities []Capability
}

func (d Device) HasCapability(needle Capability) bool {
	for _, heystack := range d.Capabilities {
		if heystack == needle {
			return true
		}
	}
	return false
}
