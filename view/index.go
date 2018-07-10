package view

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r * http.Request){
	files := GenerateFiles(Webroot+"index.html", Webroot+"include")
	t := LoadTemplate(files...)
	t.Execute(w, nil)
}