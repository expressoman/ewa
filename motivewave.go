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
func (mw *mwQuery) findImpulser(
	startPrice float64,
	endPrice float64,
) (Impulser, bool) {

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

		outImpulse := &impulse{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.Wave5),
		},
			degree: degreeFromSrting(w.Degree),
		}

		if w1, found := mw.findImpulser(w.Origin.P, w.Wave1.P); found {
			outImpulse.w1 = w1
		}

		if w2, found := mw.findCorrectioner(w.Wave1.P, w.Wave2.P); found {
			outImpulse.w2 = w2
		}

		if w3, found := mw.findImpulser(w.Wave2.P, w.Wave3.P); found {
			outImpulse.w3 = w3
		}

		if w4, found := mw.findCorrectioner(w.Wave3.P, w.Wave4.P); found {
			outImpulse.w4 = w4
		}

		if w5, found := mw.findImpulser(w.Wave4.P, w.Wave5.P); found {
			outImpulse.w5 = w5
		}

		logContext.Debugf("Found impulse %#v", outImpulse)

		return outImpulse, true
	}

	logContext.Debug("Not found")

	return nil, false
}

func (mw *mwQuery) findCorrectioner(
	startPrice float64,
	endPrice float64,
) (Correctioner, bool) {

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

			outFlat := &flat{Waver: &wave{
				base: mwPointToPoint(w.Origin),
				end:  mwPointToPoint(w.WaveA),
			},
				degree: degreeFromSrting(w.Degree),
			}

			if a, found := mw.findCorrectioner(w.Origin.P, w.WaveA.P); found {
				outFlat.a = a
			}

			if b, found := mw.findCorrectioner(w.WaveA.P, w.WaveB.P); found {
				outFlat.b = b
			}

			if c, found := mw.findImpulser(w.WaveB.P, w.WaveC.P); found {
				outFlat.c = c
			}

			logContext.Debugf("Found flat %#v", outFlat)

			return outFlat, true
		}

		// Zigzag

		zigzag := &zigzag{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveA),
		},
			degree: degreeFromSrting(w.Degree),
		}

		if a, found := mw.findImpulser(w.Origin.P, w.WaveA.P); found {
			zigzag.a = a
		}

		if b, found := mw.findCorrectioner(w.WaveA.P, w.WaveB.P); found {
			zigzag.b = b
		}

		if c, found := mw.findImpulser(w.WaveB.P, w.WaveC.P); found {
			zigzag.c = c
		}

		logContext.Debugf("Found zigzag %#v", zigzag)

		return zigzag, true
	}

	// Look in all triangles
	for _, w := range mw.Triangles {

		// Check if beginning and end of the wave are correct
		if w.Origin.P != startPrice || w.WaveE.P != endPrice {
			continue
		}

		triangle := &triangle{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveE),
		},
			degree: degreeFromSrting(w.Degree),
		}

		if a, found := mw.findCorrectioner(w.Origin.P, w.WaveA.P); found {
			triangle.a = a
		}

		if b, found := mw.findCorrectioner(w.WaveA.P, w.WaveB.P); found {
			triangle.b = b
		}

		if c, found := mw.findCorrectioner(w.WaveB.P, w.WaveC.P); found {
			triangle.c = c
		}

		if d, found := mw.findCorrectioner(w.WaveC.P, w.WaveD.P); found {
			triangle.d = d
		}

		if e, found := mw.findCorrectioner(w.WaveD.P, w.WaveE.P); found {
			triangle.e = e
		}

		logContext.Debugf("Found triangle %v", triangle)

		return triangle, true
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

			outTripple := &triple{Waver: &wave{
				base: mwPointToPoint(w.Origin),
				end:  mwPointToPoint(w.WaveZ),
			}}

			if w, found := mw.findCorrectioner(w.Origin.P, w.WaveW.P); found {
				outTripple.w = w
			}

			if x, found := mw.findCorrectioner(w.WaveW.P, w.WaveX.P); found {
				outTripple.x = x
			}

			if y, found := mw.findCorrectioner(w.WaveX.P, w.WaveY.P); found {
				outTripple.y = y
			}

			if x2, found := mw.findCorrectioner(w.WaveY.P, w.WaveX2.P); found {
				outTripple.x2 = x2
			}

			if z, found := mw.findCorrectioner(w.WaveX2.P, w.WaveZ.P); found {
				outTripple.z = z
			}

			logContext.Debugf("Found triple %v", outTripple)

			return outTripple, true
		}

		// Combo

		outCombo := &combo{Waver: &wave{
			base: mwPointToPoint(w.Origin),
			end:  mwPointToPoint(w.WaveY),
		},
			degree: degreeFromSrting(w.Degree),
		}

		if w, found := mw.findCorrectioner(w.Origin.P, w.WaveW.P); found {
			outCombo.w = w
		}

		if x, found := mw.findCorrectioner(w.WaveW.P, w.WaveX.P); found {
			outCombo.x = x
		}

		if y, found := mw.findCorrectioner(w.WaveX.P, w.WaveY.P); found {
			outCombo.y = y
		}

		logContext.Debugf("Found combo %v", outCombo)

		return outCombo, true
	}

	logContext.Debug("Not found")

	return nil, false
}

func (mw *mwQuery) markup() (*Markup, error) {
	return &Markup{}, nil
}
