package ewa

func newImpulseShort(start, end *point, degree DegreeType) Impulser {
	return &impulse{
		Waver: &wave{
			base: start,
			end:  end,
		},
		degree: degree,
	}
}

func newCorrectionShort(base *point, end *point, degree DegreeType) Correctioner {
	return &correction{Waver: &wave{base: base, end: end}, degree: degree}
}

func newImpulse(points [6]*point, degree DegreeType) Impulser {
	subDegree := degree - 1
	return &impulse{
		Waver:  &wave{base: points[0], end: points[5]},
		degree: degree,
		w1:     newImpulseShort(points[0], points[1], subDegree),
		w2:     newCorrectionShort(points[1], points[2], subDegree),
		w3:     newImpulseShort(points[2], points[3], subDegree),
		w4:     newCorrectionShort(points[3], points[4], subDegree),
		w5:     newImpulseShort(points[4], points[5], subDegree),
	}
}
