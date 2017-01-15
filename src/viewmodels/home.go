package viewmodels

import ()

type Home struct {
	Head       Head
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

func GetHome() Home {
	result := Home{
		Head: Head{
			Title:       "Weerwolven van Wakkerdam",
			Description: "Speluitleg van Weerwolven van Wakkerdam",
			Keywords:    "weerwolven, weerwolven van wakkerdam, witte weerwolf, weerwolf, het dorp, volle maan, het pact, speluitleg, rollenspel, uitbereiding",
		},
	}
	return result
}
