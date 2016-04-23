package ewa

import (
	"os"
	"time"

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

	FIt("findImpulser", func() {
		mw := &mwQuery{}

		err := mw.importMotiveWaveXML(path)
		Expect(err).Should(Succeed())

		impulserX := &impulse{Waver: &wave{
			base: &point{p: 2106.8, t: time.Now()},
			end:  &point{p: 1806.3, t: time.Now()},
		},
			degree: Primary,
		}

		mw.findImpulser(2106.8, 1806.3, impulserX)
	})
})
