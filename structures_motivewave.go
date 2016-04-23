package ewa

type mwQuery struct {
	Impulses        []mwImpulse     `xml:"graph>impulse"`
	ImpulsesLeading []mwImpulse     `xml:"graph>leading_diagonal"`
	ImpulsesEnding  []mwImpulse     `xml:"graph>ending_diagonal"`
	Corrections     []mwCorrection  `xml:"graph>correction"`
	Triangles       []mwTriangle    `xml:"graph>triangle"`
	TripleCombo     []mwComboTriple `xml:"graph>triple_combo"`
}

type mwWave struct {
	Time  int64   `xml:"time,attr"`
	Value float64 `xml:"value,attr"`
}

type mwImpulse struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwWave `xml:"origin"`
	Wave1  mwWave `xml:"wave1"`
	Wave2  mwWave `xml:"wave2"`
	Wave3  mwWave `xml:"wave3"`
	Wave4  mwWave `xml:"wave4"`
	Wave5  mwWave `xml:"wave5"`
}

type mwCorrection struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwWave `xml:"origin"`
	WaveA  mwWave `xml:"waveA"`
	WaveB  mwWave `xml:"waveB"`
	WaveC  mwWave `xml:"waveC"`
}

type mwComboTriple struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwWave `xml:"origin"`
	WaveW  mwWave `xml:"waveW"`
	WaveX  mwWave `xml:"waveX"`
	WaveY  mwWave `xml:"waveY"`
	WaveX2 mwWave `xml:"waveX2"`
	WaveZ  mwWave `xml:"waveZ"`
}

type mwTriangle struct {
	ID       int64  `xml:"id,attr"`
	ParentID int64  `xml:"parentId,attr"`
	Degree   string `xml:"degree,attr"`

	Origin mwWave `xml:"origin"`
	WaveA  mwWave `xml:"waveA"`
	WaveB  mwWave `xml:"waveB"`
	WaveC  mwWave `xml:"waveC"`
	WaveD  mwWave `xml:"waveD"`
	WaveE  mwWave `xml:"waveE"`
	WaveF  mwWave `xml:"waveF"`
	WaveG  mwWave `xml:"waveG"`
	WaveH  mwWave `xml:"waveH"`
	WaveI  mwWave `xml:"waveI"`
}
