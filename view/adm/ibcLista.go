package adm

import (
	"net/http"
	"ibcWiki/controller/adm"
	"ibcWiki/view"
	"ibcWiki/view/ibc"
	"github.com/gorilla/mux"
	
)

type listaPagina struct{
	Serie string 
	Ibcs []string
}

func IBCListaHandler(w http.ResponseWriter, r *http.Request){
	var pagina listaPagina
	files := view.GenerateFiles(view.ADMRoot+"ibc-lista.html", view.ADMRoot+"include")
	logado := adm.VerificarSessao(r)
	if logado {
		vars := mux.Vars(r)
		index := vars["serie"]
		//Listar os ibcs de suas respectivas séries
		bixos := ibc.PegarLista(index)
		t := view.LoadTemplate(files...)
		pagina = listaPagina{Serie: index, Ibcs: bixos}
		t.Execute(w, &pagina)
	} else {
		http.Redirect(w, r, "/mod/", 302)
	}
}