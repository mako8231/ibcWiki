package adm

import (
	"net/http"
	"ibcWiki/view"
	"ibcWiki/controller/adm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	logado := adm.VerificarSessao(r)
	if logado{
		http.Redirect(w, r, "/mod/painel", 302)
	} else {
		files := view.GenerateFiles(view.ADMRoot+"login.html", view.ADMRoot+"include")
		t := view.LoadTemplate(files...)
		t.Execute(w, nil)
	}

}

func VerificarCredenciaisHandler(w http.ResponseWriter, r *http.Request){
	sucesso := adm.VerificarCredenciais(r, w)
	if !sucesso{
		http.Redirect(w, r, "/mod/", 302)
		return 
	} 

	http.Redirect(w, r, "/mod/painel/", 302)
}