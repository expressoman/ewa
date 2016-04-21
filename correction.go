package ewa

func newCorrectionShort(base *point, end *point, degree DegreeType) *correction {
	return &correction{wave: &wave{base: base, end: end}, degree: degree}
}
