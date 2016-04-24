package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
		"/OANDA/analysis/CFD/SPX500USD" +
		"/Primary Analysis.mwml"

	It("Import", func() {
		mw := &mwQuery{}

		err := mw.importMotiveWaveXML(path)
		Expect(err).Should(Succeed())
	})

	FIt("Parse", func() {
		mw := &mwQuery{}

		err := mw.importMotiveWaveXML(path)
		Expect(err).Should(Succeed())

		_, err = mw.parse()
		Expect(err).Should(Succeed())
	})
})
