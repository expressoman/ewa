package ewa

func (z zigzag) HasSub() bool {
	return z.a != nil && z.b != nil && z.c != nil
}

func (z zigzag) Degree() DegreeType {
	return z.degree
}

func (z zigzag) Type() CorrectionType {
	return ZigZag
}
