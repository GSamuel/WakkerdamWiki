package viewmodels

import ()

type Home struct {
	Title      string
	Characters []Character
}

type Character struct {
	Id               int
	ImageUrl         string
	Name             string
	DescriptionLines []string
}

func GetHome() Home {
	result := Home{
		Title: "Weerwolven van Wakkerdam",
	}
	return result
}
