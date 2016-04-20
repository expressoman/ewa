package ewa

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWave(t *testing.T) {

	Convey("Wave", t, func() {
		now := time.Now()
		wave := &Wave{
			Base: &Point{T: now, P: 1890.32},
			End:  &Point{T: now.Add(-time.Hour * 16), P: 2100.54},
		}

		wave2 := &Wave{
			Base: &Point{T: now, P: 210054},
			End:  &Point{T: now.Add(-time.Hour * 16), P: 1890.32},
		}

		Convey("Interface", func() {
			Convey("Up", func() {
				So(wave.Up(), ShouldBeTrue)
			})

			Convey("Fib", func() {
				So(wave.Fib(0.5), ShouldEqual, 1995.43)
				So(wave2.Fib(0.3), ShouldEqual, 1953.38)
			})

			Convey("Len", func() {
				So(wave.Len(), ShouldEqual, 210.22)
			})

			Convey("Time", func() {
				So(wave.Time(), ShouldEqual, time.Duration(-time.Hour*16))
			})

			Convey("Begins", func() {
				So(wave.Begins().Equal(wave.Base.T), ShouldBeTrue)
			})

			Convey("Ends", func() {
				So(wave.Ends().Equal(wave.End.T), ShouldBeTrue)
			})

			Convey("Starts", func() {
				So(wave.Starts(), ShouldEqual, wave.Base.P)
			})

			Convey("Tops", func() {
				So(wave.Tops(), ShouldEqual, wave.End.P)
			})
		})
	})
}
