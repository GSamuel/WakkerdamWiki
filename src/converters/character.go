package converters

import (
	"../models"
	"../viewmodels"
)

func ConvertCharacterToViewModel(character models.Character) viewmodels.Character {
	result := viewmodels.Character{
		Id:               character.Id(),
		Name:             character.Name(),
		ImageUrl:         character.ImageUrl(),
		DescriptionLines: character.DescriptionLines(),
	}

	return result
}
