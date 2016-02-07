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
	Id      uint16
	Parent  uint16
	Impulse bool
	Degree  uint8
	Base    Point
	End     Point
}

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
