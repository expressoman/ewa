package ewa

func (f flat) HasSub() bool {
	return f.a != nil && f.b != nil && f.c != nil
}

func (f flat) Degree() DegreeType {
	return f.degree
}

func (f flat) Type() CorrectionType {
	return Flat
}
