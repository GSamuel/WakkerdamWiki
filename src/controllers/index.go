package controllers

import (
	"../controllers/util"
	"../viewmodels"
	"net/http"
	"text/template"
)

type indexController struct {
	template *template.Template
}

func (this *indexController) get(w http.ResponseWriter, req *http.Request) {

	vm := viewmodels.GetIndex()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}
