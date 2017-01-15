package controllers

import (
	"../controllers/util"
	"../converters"
	"../models"
	"../viewmodels"
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
	if err == nil {
		characters := models.GetCharacters()
		var characterVM viewmodels.Character

		for _, character := range characters {
			if character.Id() == id {
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
	} else {
		responseWriter.WriteHeader(404)
	}
}
