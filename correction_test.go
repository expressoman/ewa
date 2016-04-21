package ewa

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Correction", func() {

	f := func(tmp Correctioner) {}
	f(zigzag{})
	f(flat{})
	f(triangle{})
	f(combo{})
	f(triple{})
})
