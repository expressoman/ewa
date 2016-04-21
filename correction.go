package ewa

func (c correction) Degree() DegreeType {
	return c.degree
}

func (c correction) HasSub() bool {
	return false
}

func (c correction) Type() CorrectionType {
	return Unknown
}
