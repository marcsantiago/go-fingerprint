package fingerprint

// Stringer is a type alias for the string primitive that allows normal strings to satisfy the fmt.Stringer interface
type Stringer string

func (s Stringer) String() string { return string(s) }
