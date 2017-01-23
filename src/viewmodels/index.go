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
			Description: "Website beschrijving",
			Keywords:    keywords,
		},
		Active: "index",
	}
	return result
}
