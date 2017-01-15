package viewmodels

import (
	"fmt"
)

type Detail struct {
	Head      Head
	Character Character
}

func GetDetail(character Character) Detail {

	description := fmt.Sprintf("Uitleg van %s", character.Name)

	if len(character.DescriptionLines) > 0 {
		description = character.DescriptionLines[0]
	}

	keywords := fmt.Sprintf("%s, weerwolven, weerwolven van wakkerdam, weerwolf, het dorp, volle maan, het pact, speluitleg, rollenspel, uitbereiding", character.Name)

	result := Detail{
		Head: Head{
			Title:       character.Name,
			Description: description,
			Keywords:    keywords,
		},
		Character: character,
	}
	return result
}
