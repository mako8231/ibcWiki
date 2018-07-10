package view

import (
	"net/http"

)

func SobreNosHandler(w http.ResponseWriter, r *http.Request){
	files := GenerateFiles(Webroot+"sobre-nos.html", Webroot+"include")
	t := LoadTemplate(files...)
	t.Execute(w, nil)
}