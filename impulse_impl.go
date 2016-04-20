package ewa

import "time"

//NewImpulseShort - creates new impulse - short notion
func NewImpulseShort(start, end *point, degree Degree) *impulse {
	return &impulse{&wave: {
		Base: base,
	},
		W5:     w5,
		Degree: degree,
	}
}

func NewImpulse(points [6]*point, degree Degree) *impulse {
	return &impulse{
		base:   points[0],
		w1:     points[1],
		w2:     points[2],
		w3:     points[3],
		w4:     points[4],
		w5:     points[5],
		degree: degree}
}

func (i impulse) W1() Wave {
	return Wave{Base: i.base, End: i.w1}
}

func (i impulse) W2() Wave {
	return Wave{Base: i.w1, End: i.w2}
}

func (i impulse) W3() Wave {
	return Wave{Base: i.w2, End: i.w3}
}

func (i impulse) W4() Wave {
	return Wave{Base: i.w3, End: i.w4}
}

func (i impulse) W5() Wave {
	return Wave{Base: i.w4, End: i.w5}
}

// Wave interface

func (i impulse) Time() time.Duration {
	return i.w5.T.Sub(i.base.T)
}

func (i impulse) Len() uint32 {
	if i.Up() {
		return i.w5.P - i.base.P
	}

	return i.base.P - i.w5.P
}

func (i impulse) Up() bool {
	return i.w5.P > i.base.P
}

func (i impulse) Retrace(level float64) uint32 {
	if i.Up() {
		return i.base.P + uint32(float64(i.Len())*level)
	}

	return i.w5.P + uint32(float64(i.Len())*level)
}

//Impulser interface
type Impulser interface {
	//W1 - ref to waves making impulse formation
	W1() *Wave
	W2() *Wave
	W3() *Wave
	W4() *Wave
	W5() *Wave

	//Impulse rules
	Rule1() bool
	Rule2() bool
	Rule3() bool

	//Check - Rule1() && Rule2() && Rule3()
	Check() bool

	//Failed - price of the W5 does not go further W5
	Failed() bool

	//Extended - Len() of this wave is the biggest
	//among all child waves of the parent wave
	Extended() bool

	//Diagonal - Wave 4 overlaps Wave 1
	Diagonal() bool
}

func (i impulse) Failed() bool {
	return i.W3().End.P >= i.W5().End.P
}

func (i impulse) Extended() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}

func (i impulse) Diagonal() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}

func (i impulse) Rule1() bool {
	return i.W1().Len() >= i.W2().Len()
}

func (i impulse) Rule2() bool {
	if i.W1().Up() {
		return i.W1().End.P < i.W4().End.P
	}

	return i.W1().End.P > i.W4().End.P
}

func (i impulse) Rule3() bool {
	return i.W3().Len() >= i.W1().Len() || i.W3().Len() >= i.W5().Len()
}

func (i impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3()
}
