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

	w1 := &Impulse{Wave: &Wave{Base: Point{T: shift(5), P: 0}, End: Point{T: shift(4), P: 2}}}
	w2 := &Correction{Wave: &Wave{Base: Point{T: shift(4), P: 2}, End: Point{T: shift(3), P: 1}}}
	w3 := &Impulse{Wave: &Wave{Base: Point{T: shift(3), P: 1}, End: Point{T: shift(2), P: 5}}}
	w4 := &Correction{Wave: &Wave{Base: Point{T: shift(2), P: 5}, End: Point{T: shift(1), P: 4}}}
	w5 := &Impulse{Wave: &Wave{Base: Point{T: shift(1), P: 5}, End: Point{T: shift(0), P: 7}}}

	i1 := Impulse{W1: w1, W2: w2, W3: w3, W4: w4, W5: w5}

	Convey("Rules and Check", t, func() {
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

	//
	//Convey("Fib", t, func() {
	//	So(wave.Fib(0.5), ShouldEqual, 199543)
	//	So(wave2.Fib(0.3), ShouldEqual, 195338)
	//})
	//
	//Convey("Len", t, func() {
	//	So(wave.Len(), ShouldEqual, 21022)
	//})
}
