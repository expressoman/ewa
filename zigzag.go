package ewa

import "time"

type ZigZag struct {
	A, C *Impulse
	B    *Correction
}

// Wave interface

func (z ZigZag) Time() time.Duration {
	return z.A.base.T.Sub(z.C.w5.T)
}

func (z ZigZag) Len() uint32 {
	if z.Up() {
		return z.C.w5.P - z.A.base.P
	}

	return z.A.base.P - z.C.w5.P
}

func (z ZigZag) Up() bool {
	return z.A.base.P < z.C.w5.P
}

func (z ZigZag) Fib(level float64) uint32 {
	if z.Up() {
		return z.C.w5.P - uint32(float64(z.Len())*level)
	}

	return z.C.w5.P + uint32(float64(z.Len())*level)
}

func (z ZigZag) Begins() time.Time {
	return z.A.base.T
}

func (z ZigZag) Ends() time.Time {
	return z.C.w5.T
}

func (z ZigZag) Starts() uint32 {
	return z.A.base.P
}

func (z ZigZag) Tops() uint32 {
	return z.C.w5.P
}
