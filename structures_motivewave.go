package ewa

type mwQuery struct {
	Impulses        []mwImpulse     `xml:"graph>impulse"`
	ImpulsesLeading []mwImpulse     `xml:"graph>leading_diagonal"`
	ImpulsesEnding  []mwImpulse     `xml:"graph>ending_diagonal"`
	Corrections     []mwCorrection  `xml:"graph>correction"`
	Triangles       []mwTriangle    `xml:"graph>triangle"`
	TripleCombo     []mwComboTriple `xml:"graph>triple_combo"`
}

type mwPoint struct {
	T int64   `xml:"time,attr"`
	P float64 `xml:"value,attr"`
}

type mwImpulse struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwPoint `xml:"origin"`
	Wave1  mwPoint `xml:"wave1"`
	Wave2  mwPoint `xml:"wave2"`
	Wave3  mwPoint `xml:"wave3"`
	Wave4  mwPoint `xml:"wave4"`
	Wave5  mwPoint `xml:"wave5"`
}

type mwCorrection struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwPoint `xml:"origin"`
	WaveA  mwPoint `xml:"waveA"`
	WaveB  mwPoint `xml:"waveB"`
	WaveC  mwPoint `xml:"waveC"`
}

type mwComboTriple struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwPoint `xml:"origin"`
	WaveW  mwPoint `xml:"waveW"`
	WaveX  mwPoint `xml:"waveX"`
	WaveY  mwPoint `xml:"waveY"`
	WaveX2 mwPoint `xml:"waveX2"`
	WaveZ  mwPoint `xml:"waveZ"`
}

type mwTriangle struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwPoint `xml:"origin"`
	WaveA  mwPoint `xml:"waveA"`
	WaveB  mwPoint `xml:"waveB"`
	WaveC  mwPoint `xml:"waveC"`
	WaveD  mwPoint `xml:"waveD"`
	WaveE  mwPoint `xml:"waveE"`
	WaveF  mwPoint `xml:"waveF"`
	WaveG  mwPoint `xml:"waveG"`
	WaveH  mwPoint `xml:"waveH"`
	WaveI  mwPoint `xml:"waveI"`
}
