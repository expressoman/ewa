package ewa

// // Wave interface
//
// func (z zigzag) Time() time.Duration {
// 	return z.A.base.T.Sub(z.C.w5.T)
// }
//
// func (z zigzag) Len() float64 {
// 	if z.Up() {
// 		return z.C.w5.P - z.A.base.P
// 	}
//
// 	return z.A.base.P - z.C.w5.P
// }
//
// func (z zigzag) Up() bool {
// 	return z.A.base.P < z.C.w5.P
// }
//
// func (z zigzag) Fib(level float64) uint32 {
// 	if z.Up() {
// 		return z.C.w5.P - uint32(float64(z.Len())*level)
// 	}
//
// 	return z.C.w5.P + uint32(float64(z.Len())*level)
// }
//
// func (z zigzag) Begins() time.Time {
// 	return z.A.base.T
// }
//
// func (z zigzag) Ends() time.Time {
// 	return z.C.w5.T
// }
//
// func (z zigzag) Starts() uint32 {
// 	return z.A.base.P
// }
//
// func (z zigzag) Tops() uint32 {
// 	return z.C.w5.P
// }
