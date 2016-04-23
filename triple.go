package ewa

func (t triple) HasSub() bool {
	return t.w != nil && t.x != nil && t.y != nil && t.x2 != nil && t.z != nil
}

func (t triple) Degree() DegreeType {
	return t.degree
}

func (t triple) Type() CorrectionType {
	return Triple
}
