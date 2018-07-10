package adm

import (
	"net/http"
	"ibcWiki/controller/adm"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	logado := adm.VerificarSessao(r)
	if logado {
		adm.LimparSessao(w)
		http.Redirect(w, r, "/mod/", 302)
	} else {
		http.Redirect(w, r, "/mod/", 302)
	}
}