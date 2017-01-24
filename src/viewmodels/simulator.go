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
			Description: "Deze Simulator is een tool en hulpmiddel voor de spelleider om beter het spel te kunnen leiden. Bereken automatisch wie er in een nacht of dag sterft en zie de volledige geschiedenis van een potje in.",
			Keywords:    keywords,
			Canonical:   "http://www.witteweerwolf.nl/simulator",
		},
		Active: "simulator",
	}
	return result
}
