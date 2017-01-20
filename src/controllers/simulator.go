package controllers

import (
	"../controllers/util"
	"../viewmodels"
	"net/http"
	"text/template"
)

type simulatorController struct {
	template *template.Template
}

func (this *simulatorController) get(w http.ResponseWriter, req *http.Request) {

	vm := viewmodels.GetSimulator()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}
