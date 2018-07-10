package adm

import (
	"net/http"
	"fmt"
	"ibcWiki/view"
	"ibcWiki/controller/adm"
	"ibcWiki/controller/ibc"
)

func CadastroIBCHandler(w http.ResponseWriter, r *http.Request){
	logado := adm.VerificarSessao(r)
	if logado{
		files := view.GenerateFiles(view.ADMRoot+"ibc-add.html", view.ADMRoot+"include")
		t := view.LoadTemplate(files...)
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/mod/", 302)
	}
}

func SalvarIBCHandler(w http.ResponseWriter, r *http.Request){
	err := ibc.CadastrarIBC(r)
	if err != nil{
		fmt.Fprintf(w, err.Error())
		return 
	}
	http.Redirect(w, r, "/mod/", 302)
	
}