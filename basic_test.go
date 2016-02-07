package ewa

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestWave(t *testing.T) {

	Convey("Basic wave", t, func() {
		now := time.Now()
		wave := &Wave{
			Base: &Point{T: now, P: 189032},
			End:  &Point{T: now.Add(-time.Hour * 16), P: 210054},
		}

		wave2 := &Wave{
			Base: &Point{T: now, P: 210054},
			End:  &Point{T: now.Add(-time.Hour * 16), P: 189032},
		}

		Convey("Up or down", func() {
			So(wave.Up(), ShouldBeTrue)
		})

		Convey("Fib", func() {
			So(wave.Fib(0.5), ShouldEqual, 199543)
			So(wave2.Fib(0.3), ShouldEqual, 195338)
		})

		Convey("Len", func() {
			So(wave.Len(), ShouldEqual, 21022)
		})

	})
}
