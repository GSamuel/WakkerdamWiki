package viewmodels

import ()

type Simulator struct {
	Active string
	Head   Head
}

func GetSimulator() Index {

	keywords := ""

	result := Index{
		Head: Head{
			Title:       "WitteWeerwolf.nl",
			Description: "Website beschrijving",
			Keywords:    keywords,
		},
		Active: "index",
	}
	return result
}
