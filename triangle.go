package ewa

//import "time"

type Triangle struct {
	A, B, C, D, E *Correction
}

// Wave interface
//func (t Triangle) Time() time.Duration {
//	return i.w5.T.Sub(i.base.T)
//}
//
//func (t Triangle) Len() uint32 {
//	if i.Up() {
//		return i.w5.P - i.base.P
//	}
//
//	return i.base.P - i.w5.P
//}
//
//func (t Triangle) Up() bool {
//	return i.w5.P > i.base.P
//}
//
//func (t Triangle) Fib(level float64) uint32 {
//	if i.Up() {
//		return i.base.P + uint32(float64(i.Len())*level)
//	}
//
//	return i.w5.P + uint32(float64(i.Len())*level)
//}
