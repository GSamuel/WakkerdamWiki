package viewmodels

import ()

type Detail struct {
	Title     string
	Character Character
}

func GetDetail() Detail {
	result := Detail{
		Title: "Weerwolven van Wakkerdam",
	}
	return result
}
