package ewa

func newCorrectionShort(base *point, end *point, degree DegreeType) *correction {
	return &correction{wave: &wave{base: base, end: end}, degree: degree}
}

func (c correction) Degree() DegreeType {
	return c.degree
}

func (c correction) HasSub() bool {
	return false
}

func (c correction) Type() CorrectionType {
	return Unknown
}
