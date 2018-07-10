package adm

import (
	"net/http"
	"ibcWiki/controller/ibc"
	"ibcWiki/controller/adm"
)

func DeleteIBCHandler(w http.ResponseWriter, r *http.Request){
	logado := adm.VerificarSessao(r)
	if logado {
		err := ibc.DeleteIBC(r)
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotFound)
			return 
		}

		http.Redirect(w, r, "/mod/", 302)
	} else {
		http.Redirect(w, r, "/mod/", 302 )
	}
}