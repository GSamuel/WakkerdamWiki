package controllers

import (
	"../controllers/util"
	"../converters"
	"../models"
	"../viewmodels"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
)

type characterController struct {
	template *template.Template
}

func (this *characterController) get(w http.ResponseWriter, req *http.Request) {

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req)
	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	characters := models.GetCharacters()
	var characterVM viewmodels.Character

	for _, character := range characters {
		if err == nil && character.Id() == id {
			characterVM = converters.ConvertCharacterToViewModel(character)
		} else if character.ImageUrl() == idRaw {
			characterVM = converters.ConvertCharacterToViewModel(character)
		}
	}

	if &characterVM == nil {
		responseWriter.WriteHeader(404)
		return
	}

	vm := viewmodels.GetDetail(characterVM)

	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)

}

func (this *characterController) redirect(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idRaw := vars["id"]

	http.Redirect(w, req, fmt.Sprintf("/rollen/%s", idRaw), 301)
}
