package controllers

import (
	"../controllers/util"
	"../converters"
	"../models"
	"../viewmodels"
	"net/http"
	"text/template"
)

type charactersController struct {
	template *template.Template
}

func (this *charactersController) get(w http.ResponseWriter, req *http.Request) {

	characters := models.GetCharacters()
	charactersVM := []viewmodels.Character{}

	for _, character := range characters {
		charactersVM = append(charactersVM, converters.ConvertCharacterToViewModel(character))
	}

	vm := viewmodels.GetCharacters()
	vm.Characters = charactersVM

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}
