package ewa

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Impulse", func() {
	It("waver interface", func() {
		f := func(tmp Impulser) {}
		f(impulse{})
	})
})
