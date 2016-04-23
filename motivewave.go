package ewa

import (
	"encoding/xml"
	"io/ioutil"
)

//ImportMotiveWaveXML - importing XML
func importMotiveWaveXML(path string) (*mwQuery, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	tmp := &mwQuery{}

	err = xml.Unmarshal(data, &tmp)

	if err != nil {
		return nil, err
	}

	return tmp, nil
}
