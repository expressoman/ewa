package ewa

//import "time"

type Flat struct {
	A, B *Correction
	C    *Impulse
}

// Wave interface

//func (f Flat) Time() time.Duration {
//	return f.w5.T.Sub(f.base.T)
//}
//
//func (f Flat) Len() uint32 {
//	if f.Up() {
//		return f.w5.P - f.base.P
//	}
//
//	return f.base.P - f.w5.P
//}
//
//func (f Flat) Up() bool {
//	return f.w5.P > f.base.P
//}
//
//func (f Flat) Fib(level float64) uint32 {
//	if f.Up() {
//		return f.base.P + uint32(float64(f.Len())*level)
//	}
//
//	return f.w5.P + uint32(float64(f.Len())*level)
//}
