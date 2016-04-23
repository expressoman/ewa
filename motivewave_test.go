package ewa

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Motivewave", func() {
	It("Import", func() {

		path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
			"/OANDA/analysis/CFD/SPX500USD" +
			"/Primary Analysis.mwml"

		_, err := importMotiveWaveXML(path)
		Expect(err).Should(Succeed())
	})
})
