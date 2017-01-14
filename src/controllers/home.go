package controllers

import (
	"../controllers/util"
	"../converters"
	"../models"
	"../viewmodels"
	"net/http"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {

	characters := models.GetCharacters()
	charactersVM := []viewmodels.Character{}

	for _, character := range characters {
		charactersVM = append(charactersVM, converters.ConvertCharacterToViewModel(character))
	}

	vm := viewmodels.GetHome()
	vm.Characters = charactersVM

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}
