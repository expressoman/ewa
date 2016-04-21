package ewa

//NewImpulseShort - creates new impulse - short notion
func newImpulseShort(start, end *point, degree DegreeType) *impulse {
	return &impulse{
		wave: &wave{
			base: start,
			end:  end,
		},
		degree: degree,
	}
}

//NewImpulse creates impulse
func newImpulse(points [6]*point, degree DegreeType) *impulse {
	subDegree := degree - 1
	return &impulse{
		wave:   &wave{base: points[0], end: points[5]},
		degree: degree,
		w1:     newImpulseShort(points[0], points[1], subDegree),
		w2:     newCorrectionShort(points[1], points[2], subDegree),
		w3:     newImpulseShort(points[2], points[3], subDegree),
		w4:     newCorrectionShort(points[3], points[4], subDegree),
		w5:     newImpulseShort(points[4], points[5], subDegree),
	}
}

func (i impulse) Degree() DegreeType {
	return i.degree
}

func (i impulse) HasSub() bool {
	return i.w1 != nil && i.w2 != nil && i.w3 != nil && i.w4 != nil && i.w5 != nil
}

// Waves

func (i impulse) W1() *impulse {
	return i.w1
}

func (i impulse) W2() *correction {
	return i.w2
}

func (i impulse) W3() *impulse {
	return i.w3
}

func (i impulse) W4() *correction {
	return i.w4
}

func (i impulse) W5() *impulse {
	return i.w5
}

// Rules

func (i impulse) Rule1() bool {
	return i.w1.Len() >= i.W2().Len()
}

func (i impulse) Rule2() bool {
	if i.w1.Up() {
		return i.w1.end.p < i.w4.end.p
	}

	return i.w1.end.p > i.w4.end.p
}

func (i impulse) Rule3() bool {
	return i.W3().Len() >= i.w1.Len() || i.W3().Len() >= i.W5().Len()
}

func (i impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3()
}

// Impulse specific

func (i impulse) Failed() bool {
	return i.w3.end.p >= i.W5().end.p
}

func (i impulse) Extended() bool {
	return i.w3.Len() >= i.w5.Len()*2 || i.w3.Len() >= i.w1.Len()*2
}

func (i impulse) Diagonal() bool {
	return i.w3.Len() >= i.w5.Len()*2 || i.w3.Len() >= i.w1.Len()*2
}
