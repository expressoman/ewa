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

// func (i impulse) W1() *impulse {
// 	return i.w1
// }
//
// func (i impulse) W2() *correction {
// 	return i.w2
// }
//
// func (i impulse) W3() *impulse {
// 	return i.w3
// }
//
// func (i impulse) W4() *correction {
// 	return i.w4
// }
//
// func (i impulse) W5() *impulse {
// 	return i.w5
// }
//
// // waver interface
//
// func (i impulse) Time() time.Duration {
// 	return i.w5.bast.Sub(i.base.t)
// }
//
// func (i impulse) Len() uint32 {
// 	if i.Up() {
// 		return i.w5.end.p - i.base.p
// 	}
//
// 	return i.base.p - i.w5.p
// }
//
// func (i impulse) Up() bool {
// 	return i.w5.P > i.base.P
// }
//
// func (i impulse) Retrace(level float64) uint32 {
// 	if i.Up() {
// 		return i.base.P + uint32(float64(i.Len())*level)
// 	}
//
// 	return i.w5.P + uint32(float64(i.Len())*level)
// }
//
// func (i impulse) Failed() bool {
// 	return i.W3().end.P >= i.W5().end.P
// }
//
// func (i impulse) Extended() bool {
// 	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
// }
//
// func (i impulse) Diagonal() bool {
// 	return i.W3().Len() >= i.W5().Len()*2 || i.W3().Len() >= i.W1().Len()*2
// }
//
// func (i impulse) Rule1() bool {
// 	return i.W1().Len() >= i.W2().Len()
// }
//
// func (i impulse) Rule2() bool {
// 	if i.W1().Up() {
// 		return i.W1().end.P < i.W4().end.P
// 	}
//
// 	return i.W1().end.P > i.W4().end.P
// }
//
// func (i impulse) Rule3() bool {
// 	return i.W3().Len() >= i.W1().Len() || i.W3().Len() >= i.W5().Len()
// }
//
// func (i impulse) Check() bool {
// 	return i.Rule1() && i.Rule2() && i.Rule3()
// }
