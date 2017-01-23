package viewmodels

import ()

type Simulator struct {
	Active string
	Head   Head
}

func GetSimulator() Index {

	keywords := "witte weerwolf, weerwolven, weerwolven van wakkerdam, simulator, spelleider, hulpmiddel, simulatie, tool"

	result := Index{
		Head: Head{
			Title:       "Simulator - WitteWeerwolf.nl",
			Description: "Website beschrijving",
			Keywords:    keywords,
		},
		Active: "simulator",
	}
	return result
}
