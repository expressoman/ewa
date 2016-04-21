package ewa

func (t triangle) HasSub() bool {
	return t.a != nil && t.b != nil && t.c != nil && t.d != nil && t.e != nil
}

func (t triangle) Degree() DegreeType {
	return t.degree
}

func (t triangle) Type() CorrectionType {
	return Triangle
}
