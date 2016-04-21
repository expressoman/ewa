package ewa

import . "github.com/onsi/ginkgo"

var _ = Describe("Wave", func() {
	It("waver interface", func() {
		f := func(tmp Waver) {}
		f(wave{})
	})
})
