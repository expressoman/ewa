package ewa

func (i impulse) Degree() DegreeType {
	return i.degree
}

func (i impulse) HasSub() bool {
	return i.W1() != nil && i.W2() != nil && i.W3() != nil && i.W4() != nil && i.W5() != nil
}

// Waves

func (i impulse) W1() Impulser {
	return i.W1()
}

func (i impulse) W2() Correctioner {
	return i.W2()
}

func (i impulse) W3() Impulser {
	return i.W3()
}

func (i impulse) W4() Correctioner {
	return i.W4()
}

func (i impulse) W5() Impulser {
	return i.W5()
}

// Rules

func (i impulse) Rule1() bool {
	return i.W1().Len() >= i.W2().Len()
}

func (i impulse) Rule2() bool {
	if i.W1().Up() {
		return i.W1().Tops() < i.W4().Tops()
	}

	return i.W1().Tops() > i.W4().Tops()
}

func (i impulse) Rule3() bool {
	return i.W3().Len() >= i.W1().Len() || i.W3().Len() >= i.W5().Len()
}

func (i impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3()
}

// Structure specific

func (i impulse) Failed() bool {
	return i.W3().Tops() >= i.W5().Tops()
}

func (i impulse) Extended() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}

func (i impulse) Diagonal() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}
