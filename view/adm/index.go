package adm

import (
	"net/http"
	"ibcWiki/view"
	"ibcWiki/controller/adm"
)

func AdmIndexHandler(w http.ResponseWriter, r *http.Request){
	files := view.GenerateFiles(view.ADMRoot+"index.html", view.ADMRoot+"include")
	logado := adm.VerificarSessao(r)
	if logado{
		t := view.LoadTemplate(files... )
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/mod/", 302)
	}

}

