package viewmodels

import (
	"fmt"
)

type Detail struct {
	Head      Head
	Character Character
	Active    string
}

func GetDetail(character Character) Detail {

	description := fmt.Sprintf("Uitleg van %s", character.Name)

	if len(character.DescriptionLines) > 0 {
		description = character.DescriptionLines[0]
	}

	keywords := fmt.Sprintf("%s, weerwolven, weerwolven van wakkerdam, witte weerwolf", character.Name)
	title := fmt.Sprintf("%s - WitteWeerwolf.nl", character.Name)

	result := Detail{
		Head: Head{
			Title:       title,
			Description: description,
			Keywords:    keywords,
		},
		Character: character,
		Active:    "rollen",
	}
	return result
}
