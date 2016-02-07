package ewa

import "time"

type Impulse struct {
	base, w1, w2, w3, w4, w5 *Point
	Degree                   uint8
}

func NewImpulse(points [6]*Point, degree uint8) *Impulse {
	return &Impulse{
		base:   points[0],
		w1:     points[1],
		w2:     points[2],
		w3:     points[3],
		w4:     points[4],
		w5:     points[5],
		Degree: degree}
}

func (i Impulse) W1() Wave {
	return Wave{Base: i.base, End: i.w1}
}

func (i Impulse) W2() Wave {
	return Wave{Base: i.w1, End: i.w2}
}

func (i Impulse) W3() Wave {
	return Wave{Base: i.w2, End: i.w3}
}

func (i Impulse) W4() Wave {
	return Wave{Base: i.w3, End: i.w4}
}

func (i Impulse) W5() Wave {
	return Wave{Base: i.w4, End: i.w5}
}

// Wave interface

func (i Impulse) Time() time.Duration {
	return i.w5.T.Sub(i.base.T)
}

func (i Impulse) Len() uint32 {
	if i.Up() {
		return i.w5.P - i.base.P
	}

	return i.base.P - i.w5.P
}

func (i Impulse) Up() bool {
	return i.w5.P > i.base.P
}

func (i Impulse) Fib(level float64) uint32 {
	if i.Up() {
		return i.base.P + uint32(float64(i.Len())*level)
	}

	return i.w5.P + uint32(float64(i.Len())*level)
}

// Impulse interface

func (i Impulse) Failed() bool {
	return i.W3().End.P >= i.W5().End.P
}

func (i Impulse) Extended() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}

func (i Impulse) Diagonal() bool {
	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
}

func (i Impulse) Rule1() bool {
	return i.W1().Len() >= i.W2().Len()
}

func (i Impulse) Rule2() bool {
	if i.W1().Up() {
		return i.W1().End.P < i.W4().End.P
	}

	return i.W1().End.P > i.W4().End.P
}

func (i Impulse) Rule3() bool {
	return i.W3().Len() >= i.W1().Len() || i.W3().Len() >= i.W5().Len()
}

func (i Impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3()
}
