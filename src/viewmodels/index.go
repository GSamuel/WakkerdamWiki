package viewmodels

import ()

type Index struct {
	Active string
	Head   Head
}

func GetIndex() Index {

	keywords := "witte weerwolf, weerwolven, weerwolven van wakkerdam, weerwolf, het dorp, volle maan, het pact, speluitleg, rollenspel, uitbereiding"

	result := Index{
		Head: Head{
			Title:       "Home - WitteWeerwolf.nl",
			Description: "De website over Weerwolven van Wakkerdam. Hier vind je een uitgebreide Speluitleg, bescrhijving van karakters en rollen en hulpmiddelen voor spelleiders.",
			Keywords:    keywords,
			Canonical:   "http://www.witteweerwolf.nl",
		},
		Active: "index",
	}
	return result
}
