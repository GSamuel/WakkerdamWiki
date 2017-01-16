package viewmodels

import ()

type Characters struct {
	Head       Head
	Active     string
	Characters []Character
}

type Character struct {
	Id               int
	ImageUrl         string
	Name             string
	DescriptionLines []string
}

type Head struct {
	Title       string
	Description string
	Keywords    string
}

func GetCharacters() Characters {
	result := Characters{
		Head: Head{
			Title:       "Rollen - WitteWeerwolf.nl",
			Description: "Speluitleg van Weerwolven van Wakkerdam",
			Keywords:    "weerwolven, weerwolven van wakkerdam, witte weerwolf, weerwolf, het dorp, volle maan, het pact, speluitleg, rollenspel, uitbereiding",
		},
		Active: "rollen",
	}
	return result
}
