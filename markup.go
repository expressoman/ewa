package ewa

import "github.com/apex/log"

//Markup struct
type Markup struct {
	Impulses    []*Impulse
	Corrections []*Correction
	Zigzags     []*Zigzag
	Flats       []*Flat
	Triangles   []*Triangle
	Combos      []*Combo
	Triples     []*Triple
}

func (m *Markup) addImpulse(impulse *Impulse) *Impulse {

	context := log.WithFields(log.Fields{
		"M": impulse.Move,
		"D": impulse.Degree,
	})

	for _, one := range m.Impulses {
		if one.Move.Base.P == impulse.Move.Base.P &&
			one.Move.End.P == impulse.Move.End.P {

			context.Debug("Impulse found")
			return one
		}
	}

	context.Debug("+Impulse")

	m.Impulses = append(m.Impulses, impulse)
	return impulse
}

func (m *Markup) addCorrection(correction *Correction) *Correction {
	context := log.WithFields(log.Fields{
		"M": correction.Move,
		"D": correction.Degree,
	})

	for _, one := range m.Corrections {
		if one.Move.Base.P == correction.Move.Base.P &&
			one.Move.End.P == correction.Move.End.P {

			context.Debug("Correction found")
			return one
		}
	}

	context.Debug("+Correction")

	m.Corrections = append(m.Corrections, correction)
	return correction
}

func (m *Markup) processImpulseWave(w mwImpulse) {

	degree := Degree(w.Degree)
	lessDegree := degree << 1

	//Generating points

	ori := PointFromMW(w.Origin)
	pW1 := PointFromMW(w.Wave1)
	pW2 := PointFromMW(w.Wave2)
	pW3 := PointFromMW(w.Wave3)
	pW4 := PointFromMW(w.Wave4)
	pW5 := PointFromMW(w.Wave5)

	//Generating sub waves

	subW1 := &Wave{Move: &Move{ori, pW1}, Degree: lessDegree}
	subW2 := &Wave{Move: &Move{pW1, pW2}, Degree: lessDegree}
	subW3 := &Wave{Move: &Move{pW2, pW3}, Degree: lessDegree}
	subW4 := &Wave{Move: &Move{pW3, pW4}, Degree: lessDegree}
	subW5 := &Wave{Move: &Move{pW4, pW5}, Degree: lessDegree}

	setImpulseWaveAdjecency([5]*Wave{subW1, subW2, subW3, subW4, subW5})

	// Generating impulse wave
	wave := &Wave{Move: &Move{ori, pW5}, Degree: degree}

	setParentWave(wave, subW1, subW2, subW3, subW4, subW5)

	// Adding primitives
	_ = m.addImpulse(&Impulse{Wave: wave})

	_ = m.addImpulse(&Impulse{Wave: subW1})
	_ = m.addImpulse(&Impulse{Wave: subW3})
	_ = m.addImpulse(&Impulse{Wave: subW5})

	_ = m.addCorrection(&Correction{Wave: subW2})
	_ = m.addCorrection(&Correction{Wave: subW4})

}

func (m *Markup) processImpulses(mwQuery *mwQuery) {
	for _, w := range mwQuery.Impulses {
		m.processImpulseWave(w)
	}

	for _, w := range mwQuery.ImpulsesLeading {
		m.processImpulseWave(w)
	}

	for _, w := range mwQuery.ImpulsesEnding {
		m.processImpulseWave(w)
	}
}

func (m *Markup) processCorrections(mwQuery *mwQuery) {

	for _, w := range mwQuery.Corrections {

		degree := Degree(w.Degree)
		lessDegree := degree << 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWA := PointFromMW(w.WaveA)
		pWB := PointFromMW(w.WaveB)
		pWC := PointFromMW(w.WaveC)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWC}, Degree: degree}

		subWA := &Wave{Move: &Move{ori, pWA}, Degree: lessDegree}
		subWB := &Wave{Move: &Move{pWA, pWB}, Degree: lessDegree}
		subWC := &Wave{Move: &Move{pWB, pWC}, Degree: lessDegree}

		setParentWave(wave, subWA, subWB, subWC)

		// Generating correction wave
		correction := &Wave{Move: &Move{ori, pWC}, Degree: degree}

		_ = m.addCorrection(&Correction{Wave: correction})

		_ = m.addCorrection(&Correction{Wave: subWB})
		_ = m.addImpulse(&Impulse{Wave: subWC})

		//TODO ZigZag or Flat
	}
}

func (m *Markup) processTriangles(mwQuery *mwQuery) {

	for _, w := range mwQuery.Triangles {

		degree := Degree(w.Degree)
		lessDegree := degree << 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWA := PointFromMW(w.WaveA)
		pWB := PointFromMW(w.WaveB)
		pWC := PointFromMW(w.WaveC)
		pWD := PointFromMW(w.WaveD)
		pWE := PointFromMW(w.WaveE)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWE}, Degree: degree}

		subWA := &Wave{Move: &Move{ori, pWA}, Degree: lessDegree}
		subWB := &Wave{Move: &Move{pWA, pWB}, Degree: lessDegree}
		subWC := &Wave{Move: &Move{pWB, pWC}, Degree: lessDegree}
		subWD := &Wave{Move: &Move{pWC, pWD}, Degree: lessDegree}
		subWE := &Wave{Move: &Move{pWD, pWE}, Degree: lessDegree}

		setParentWave(wave, subWA, subWB, subWC, subWD, subWE)

		// Generating triangle wave
		triangleWave := &Wave{Move: &Move{ori, pWE}, Degree: degree}

		triangleCorrection := m.addCorrection(&Correction{Wave: triangleWave})

		triangle := &Triangle{
			A: m.addCorrection(&Correction{Wave: subWA}),
			B: m.addCorrection(&Correction{Wave: subWB}),
			C: m.addCorrection(&Correction{Wave: subWC}),
			D: m.addCorrection(&Correction{Wave: subWD}),
			E: m.addCorrection(&Correction{Wave: subWE}),
		}

		log.WithField("Triangle", triangle).Debug("+Triangle")

		triangleCorrection.Triangle = triangle

		m.Triangles = append(m.Triangles, triangle)

		//TODO 9 wave triangle
	}
}

func (m *Markup) processTripleWave(w mwComboTriple) {

	degree := Degree(w.Degree)
	lessDegree := degree << 1

	//Generating points

	ori := PointFromMW(w.Origin)
	pWW := PointFromMW(w.WaveW)
	pWX := PointFromMW(w.WaveX)
	pWY := PointFromMW(w.WaveY)
	pWX2 := PointFromMW(w.WaveX2)
	pWZ := PointFromMW(w.WaveZ)

	//Generating sub waves

	wave := &Wave{Move: &Move{ori, pWZ}, Degree: degree}

	subWW := &Wave{Move: &Move{ori, pWW}, Degree: lessDegree}
	subWX := &Wave{Move: &Move{pWW, pWX}, Degree: lessDegree}
	subWY := &Wave{Move: &Move{pWX, pWY}, Degree: lessDegree}
	subWX2 := &Wave{Move: &Move{pWY, pWX2}, Degree: lessDegree}
	subWZ := &Wave{Move: &Move{pWX2, pWZ}, Degree: lessDegree}

	setParentWave(wave, subWW, subWX, subWY, subWX2, subWZ)

	// Generating triangle wave
	tripleWave := &Wave{Move: &Move{ori, pWZ}, Degree: degree}

	tripleCorrection := m.addCorrection(&Correction{Wave: tripleWave})

	triple := &Triple{
		W:  m.addCorrection(&Correction{Wave: subWW}),
		X:  m.addCorrection(&Correction{Wave: subWX}),
		Y:  m.addCorrection(&Correction{Wave: subWY}),
		X2: m.addCorrection(&Correction{Wave: subWX2}),
		Z:  m.addCorrection(&Correction{Wave: subWZ}),
	}

	tripleCorrection.Triple = triple

	log.WithField("Triple", triple).Debug("+Triple")

	m.Triples = append(m.Triples, triple)
}

func (m *Markup) processComboWave(w mwComboTriple) {

	degree := Degree(w.Degree)
	lessDegree := degree << 1

	//Generating points

	ori := PointFromMW(w.Origin)
	pWW := PointFromMW(w.WaveW)
	pWX := PointFromMW(w.WaveX)
	pWY := PointFromMW(w.WaveY)

	//Generating sub waves

	wave := &Wave{Move: &Move{ori, pWY}, Degree: degree}

	subWW := &Wave{Move: &Move{ori, pWW}, Degree: lessDegree}
	subWX := &Wave{Move: &Move{pWW, pWX}, Degree: lessDegree}
	subWY := &Wave{Move: &Move{pWX, pWY}, Degree: lessDegree}

	setParentWave(wave, subWW, subWX, subWY)

	// Generating triangle wave
	comboWave := &Wave{Move: &Move{ori, pWY}, Degree: degree}

	comboCorrection := m.addCorrection(&Correction{Wave: comboWave})

	combo := &Combo{
		W: m.addCorrection(&Correction{Wave: subWW}),
		X: m.addCorrection(&Correction{Wave: subWX}),
		Y: m.addCorrection(&Correction{Wave: subWY}),
	}

	comboCorrection.Combo = combo

	log.WithField("Combo", combo).Debug("+Combo")

	m.Combos = append(m.Combos, combo)
}

func (m *Markup) processTripleCombo(mwQuery *mwQuery) {
	for _, w := range mwQuery.TripleCombo {
		if w.WaveZ.T != 0 {
			m.processTripleWave(w)
		} else {
			m.processComboWave(w)
		}
	}
}
