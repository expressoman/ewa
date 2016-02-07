package ewa

type Impulse struct {
	*Wave
	W1, W3, W5 *Impulse
	W2, W4     *Correction
}

func (i Impulse) Failed() bool {
	return i.W3.End.P >= i.W5.End.P
}

func (i Impulse) Extended() bool {
	return i.W3.Len() >= i.W5.Len()*2 || i.W3.Len() >= i.W1.Len()*2
}

func (i Impulse) Rule1() bool {
	return i.W1.Len() >= i.W2.Len()
}

func (i Impulse) Rule2() bool {
	if i.W1.Up() {
		return i.W1.End.P < i.W4.End.P
	}

	return i.W1.End.P > i.W4.End.P
}

func (i Impulse) Rule3() bool {
	return i.W3.Len() >= i.W1.Len() || i.W3.Len() >= i.W5.Len()
}

func (i Impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3() && !i.W2.Impulse && !i.W4.Impulse
}
