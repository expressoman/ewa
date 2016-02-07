package ewa

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var _ = fmt.Print

func TestImpulse(t *testing.T) {
	now := time.Now()

	shift := func(s int) time.Time {
		return now.Add(-time.Hour * time.Duration(s))
	}

	points := [...]*Point{
		&Point{T: shift(5), P: 10},
		&Point{T: shift(4), P: 12},
		&Point{T: shift(3), P: 11},
		&Point{T: shift(2), P: 15},
		&Point{T: shift(1), P: 14},
		&Point{T: shift(0), P: 17},
	}

	i1 := NewImpulse(points, 1)

	Convey("Impulse", t, func() {
		Convey("Rules and Check", func() {

			Convey("Rule1", func() {
				So(i1.Rule1(), ShouldBeTrue)
			})

			Convey("Rule2", func() {
				So(i1.Rule2(), ShouldBeTrue)
			})

			Convey("Rule3", func() {
				So(i1.Rule3(), ShouldBeTrue)
			})

			Convey("Check", func() {
				So(i1.Check(), ShouldBeTrue)
			})
		})

		Convey("Fib and Len", func() {
			wave1 := &Wave{
				Base: &Point{T: shift(5), P: 1800},
				End:  &Point{T: shift(3), P: 2000}}

			wave2 := &Wave{
				Base: &Point{T: shift(5), P: 2100},
				End:  &Point{T: shift(3), P: 1800}}

			Convey("Fib", func() {
				So(wave1.Fib(0.5), ShouldEqual, 1900)
				So(wave2.Fib(0.33), ShouldEqual, 1899)
			})

			Convey("Len", func() {
				So(wave1.Len(), ShouldEqual, 200)
			})
		})
	})
}
