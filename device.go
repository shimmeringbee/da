package da

type Device struct {
	Gateway      Gateway
	Identifier   Identifier
	Capabilities []Capability
}

func (d Device) HasCapability(needle Capability) bool {
	for _, haystack := range d.Capabilities {
		if haystack == needle {
			return true
		}
	}
	return false
}
