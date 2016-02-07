package ewa

import (
	"time"
)

//const (
//	_             = iota
//	primary uint8 = iota
//	intermediate
//	minor
//	minute
//	minuette
//	subminuette
//	micro
//	submicro
//	miniscule
//	nano
//	subnano
//	pico
//)

type Point struct {
	P uint32
	T time.Time
}

type Wave struct {
	Base *Point
	End  *Point
}

// Wave interface

func (w Wave) Time() time.Duration {
	return w.End.T.Sub(w.Base.T)
}

func (w Wave) Len() uint32 {
	if w.Up() {
		return w.End.P - w.Base.P
	}

	return w.Base.P - w.End.P
}

func (w Wave) Up() bool {
	return w.End.P > w.Base.P
}

func (w Wave) Fib(level float64) uint32 {
	if w.Up() {
		return w.Base.P + uint32(float64(w.Len())*level)
	}

	return w.End.P + uint32(float64(w.Len())*level)
}

func (w Wave) Begins() time.Time {
	return w.Base.T
}

func (w Wave) Ends() time.Time {
	return w.End.T
}

func (w Wave) Starts() uint32 {
	return w.Base.P
}

func (w Wave) Tops() uint32 {
	return w.End.P
}
