package ewa

import (
	"encoding/xml"
	"io/ioutil"
	"math"
	"time"

	"github.com/apex/log"
)

func (mw *mwQuery) importMotiveWaveXML(path string) error {

	log.Debug("Importing " + path)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return xml.Unmarshal(data, mw)
}

func degreeFromSrting(degree string) DegreeType {
	switch degree {
	case "Primary":
		return Primary
	case "Intermediate":
		return Intermediate
	case "Minor":
		return Minor
	case "Minute":
		return Minute
	case "Minuette":
		return Minuette
	case "Subminuette":
		return Subminuette
	case "Micro":
		return Micro
	case "Submicro":
		return Submicro
	case "Miniscule":
		return Miniscule
	case "Nano":
		return Nano
	case "Subnano":
		return Subnano
	case "Pico":
		return Pico
	}

	return Minuette
}

func mwPointToPoint(mwp mwPoint) *point {
	return &point{t: time.Unix(mwp.T, 0), p: mwp.P}
}

//TODO refactor to pointer
func (mw *mwQuery) findImpulser(startPrice float64, endPrice float64, refWave Impulser) {

	logContext := log.WithFields(&log.Fields{
		"startPrice": startPrice,
		"endPrice":   endPrice,
	})

	logContext.Debug("Looking for impulser")

	//TODO Leading Ending
	for _, w := range mw.Impulses {

		// Check if beginning and end of the wave are correct
		if w.Origin.P != startPrice || w.Wave5.P != endPrice {
			continue
		}

		refWave := &impulse{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.Wave5),
		},
			degree: degreeFromSrting(w.Degree),
		}

		logContext.Debugf("Found impulse %#v", refWave)

		mw.findImpulser(w.Origin.P, w.Wave1.P, refWave.w1)
		mw.findCorrectioner(w.Wave1.P, w.Wave2.P, refWave.w2)
		mw.findImpulser(w.Wave2.P, w.Wave3.P, refWave.w3)
		mw.findCorrectioner(w.Wave3.P, w.Wave4.P, refWave.w4)
		mw.findImpulser(w.Wave4.P, w.Wave5.P, refWave.w5)

		return
	}

	logContext.Debug("Not found")
}

func (mw *mwQuery) findCorrectioner(startPrice float64, endPrice float64, refWave Correctioner) {

	logContext := log.WithFields(&log.Fields{
		"startPrice": startPrice,
		"endPrice":   endPrice,
	})

	logContext.Debugf("Looking for correctioner")

	// Look in all corrections
	for _, w := range mw.Corrections {

		// Check if beginning and end of the wave are correct
		if w.Origin.P != startPrice || w.WaveC.P != endPrice {
			continue
		}

		// Check for flat

		if math.Mod(w.Origin.P, w.WaveA.P)*.9 <= math.Mod(w.WaveA.P, w.WaveB.P) {

			refWave := &flat{Waver: &wave{
				base: mwPointToPoint(w.Origin),
				end:  mwPointToPoint(w.WaveA),
			},
				degree: degreeFromSrting(w.Degree),
			}

			logContext.Debugf("Found flat %#v", refWave)

			mw.findCorrectioner(w.Origin.P, w.WaveA.P, refWave.a)
			mw.findCorrectioner(w.WaveA.P, w.WaveB.P, refWave.b)
			mw.findImpulser(w.WaveB.P, w.WaveC.P, refWave.c)

			return
		}

		// Zigzag

		refWave := &zigzag{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveA),
		},
			degree: degreeFromSrting(w.Degree),
		}

		logContext.Debugf("Found zigzag %#v", refWave)

		mw.findImpulser(w.Origin.P, w.WaveA.P, refWave.a)
		mw.findCorrectioner(w.WaveA.P, w.WaveB.P, refWave.b)
		mw.findImpulser(w.WaveB.P, w.WaveC.P, refWave.c)

		return
	}

	// Look in all triangles
	for _, w := range mw.Triangles {

		// Check if beginning and end of the wave are correct
		if w.Origin.P != startPrice || w.WaveE.P != endPrice {
			continue
		}

		refWave := &triangle{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveE),
		},
			degree: degreeFromSrting(w.Degree),
		}

		logContext.Debugf("Found triangle %v", refWave)

		mw.findCorrectioner(w.Origin.P, w.WaveA.P, refWave.a)
		mw.findCorrectioner(w.WaveA.P, w.WaveB.P, refWave.b)
		mw.findCorrectioner(w.WaveB.P, w.WaveC.P, refWave.c)
		mw.findCorrectioner(w.WaveC.P, w.WaveD.P, refWave.d)
		mw.findCorrectioner(w.WaveD.P, w.WaveE.P, refWave.e)

		return
	}

	// Look in all triple_combo
	for _, w := range mw.TripleCombo {

		// Check if beginning
		if w.Origin.P != startPrice {
			continue
		}

		// Check if Z exists
		if w.WaveZ.P != 0 && w.WaveZ.P != endPrice {
			continue
		}

		if w.WaveY.P != endPrice {
			continue
		}

		if w.WaveZ.P != 0 {

			refWave := &triple{Waver: &wave{
				base: mwPointToPoint(w.Origin),
				end:  mwPointToPoint(w.WaveZ),
			}}

			logContext.Debugf("Found triple %v", refWave)

			mw.findCorrectioner(w.Origin.P, w.WaveW.P, refWave.w)
			mw.findCorrectioner(w.WaveW.P, w.WaveX.P, refWave.x)
			mw.findCorrectioner(w.WaveX.P, w.WaveY.P, refWave.y)
			mw.findCorrectioner(w.WaveY.P, w.WaveX2.P, refWave.x2)
			mw.findCorrectioner(w.WaveX2.P, w.WaveZ.P, refWave.z)

			return
		}

		// Combo

		refWave := &combo{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveY),
		},
			degree: degreeFromSrting(w.Degree),
		}

		logContext.Debugf("Found combo %v", refWave)

		mw.findCorrectioner(w.Origin.P, w.WaveW.P, refWave.w)
		mw.findCorrectioner(w.WaveW.P, w.WaveX.P, refWave.x)
		mw.findCorrectioner(w.WaveX.P, w.WaveY.P, refWave.y)

		return
	}

	logContext.Debug("Not found")
}

func (mw *mwQuery) markup() (*Markup, error) {
	return &Markup{}, nil
}
