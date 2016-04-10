package ewa

//import "time"

type Combo struct {
	W, X, Y *Correction
}

// Wave interface

//func (c Combo) Time() time.Duration {
//	return c.W.T.Sub(i.base.T)
//}
//
//func (c Combo) Len() uint32 {
//	if i.Up() {
//		return i.w5.P - i.base.P
//	}
//
//	return i.base.P - i.w5.P
//}
//
//func (c Combo) Up() bool {
//	return i.w5.P > i.base.P
//}
//
//func (c Combo) Fib(level float64) uint32 {
//	if i.Up() {
//		return i.base.P + uint32(float64(i.Len())*level)
//	}
//
//	return i.w5.P + uint32(float64(i.Len())*level)
//}
