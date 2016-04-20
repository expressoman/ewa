package test_impulse

import (
	"testing"
	"time"
)

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

	i1 := FullImpulse(points, 1)

	Convey("Impulse", t, func() {
		Convey("Impulse interface", func() {

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

			Convey("Failed", func() {

				Convey("Is", func() {
					points2 := [...]*Point{
						&Point{T: shift(5), P: 10},
						&Point{T: shift(4), P: 12},
						&Point{T: shift(3), P: 11},
						&Point{T: shift(2), P: 15},
						&Point{T: shift(1), P: 14},
						&Point{T: shift(0), P: 15},
					}

					failedImpulse := FullImpulse(points2, 1)

					So(failedImpulse.Failed(), ShouldBeTrue)
				})

				Convey("Not", func() {
					So(i1.Failed(), ShouldBeFalse)
				})
			})

			Convey("Extended", func() {

				Convey("Is", func() {
					So(i1.Extended(), ShouldBeTrue)
				})

				Convey("Not", func() {
					points2 := [...]*Point{
						&Point{T: shift(5), P: 10},
						&Point{T: shift(4), P: 12},
						&Point{T: shift(3), P: 11},
						&Point{T: shift(2), P: 14},
						&Point{T: shift(1), P: 13},
						&Point{T: shift(0), P: 15},
					}
					notExtImpulse := FullImpulse(points2, 1)
					So(notExtImpulse.Extended(), ShouldBeFalse)
				})
			})

			Convey("Diagonal", func() {
				So(i1.Diagonal(), ShouldBeTrue)
			})
		})

		Convey("Waver interface", func() {

			points1 := [...]*Point{
				&Point{T: shift(5), P: 100},
				&Point{T: shift(4), P: 120},
				&Point{T: shift(3), P: 110},
				&Point{T: shift(2), P: 150},
				&Point{T: shift(1), P: 140},
				&Point{T: shift(0), P: 170},
			}

			i1 := FullImpulse(points1, 1)

			Convey("Fib", func() {
				So(i1.Fib(0.5), ShouldEqual, 135)
				So(i1.Fib(0.33), ShouldEqual, 123)
			})

			Convey("Len", func() {
				So(i1.Len(), ShouldEqual, 70)
			})
		})
	})
}
