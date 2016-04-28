package ewa

import (
	"encoding/xml"
	"io/ioutil"
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

//PointFromMW - Point from mwPoint
func PointFromMW(mwp mwPoint) *Point {
	return &Point{T: time.Unix(mwp.T, 0), P: mwp.P}
}

func setParentWave(parent *Wave, children ...*Wave) {
	for _, one := range children {
		one.Parent = parent
	}
}

func setImpulseWaveAdjecency(waves [5]*Wave) {
	waves[0].Next = waves[1]

	waves[1].Prev, waves[1].Next = waves[0], waves[2]
	waves[2].Prev, waves[2].Next = waves[1], waves[3]
	waves[3].Prev, waves[3].Next = waves[2], waves[4]

	waves[4].Prev = waves[3]
}

func (mw *mwQuery) parse() (*Markup, error) {
	markup := &Markup{}

	markup.processImpulses(mw)
	markup.processCorrections(mw)
	markup.processTriangles(mw)
	markup.processTripleCombo(mw)

	markup.printStack()

	return markup, nil
}
