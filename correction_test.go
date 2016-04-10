package ewa

import (
	"fmt"
	//. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var _ = fmt.Print
var _ = time.April

func TestCorrection(t *testing.T) {
	//now := time.Now()
	//
	//shift := func(s int) time.Time {
	//	return now.Add(-time.Hour * time.Duration(s))
	//}
	//
	//pBase := &Point{2000, shift(5)}
	//pA := &Point{1800, shift(4)}
	//pB := &Point{1980, shift(3)}
	//pC := &Point{1700, shift(2)}
	//
	//wA := &ZigZag{A: &Correction{&Wave{Base: pBase, End: pA}}}

	c := FullCorrection(ZigZag{}, 1)

	fmt.Printf("%#v", c)
}
