package ibc

import (
	"github.com/gorilla/mux"
	"net/http"
	"ibcWiki/controller/ibc"
	"ibcWiki/view"
	"html/template"
	"io/ioutil"
)

func PaginaIBCHandler(w http.ResponseWriter, r *http.Request){
	//files := view.GenerateFiles(view.Webroot+"ibc-view.html", "include")
	//carregar o objeto principal
	b, err := ioutil.ReadFile(view.Webroot+"ibc-view.html")
	if err != nil{
		http.Error(w, err.Error(), 404)
		return 
	}
	t,_ := template.Must(template.New("ibc").Funcs(view.Funcs).Parse(string(b))).ParseFiles(view.Webroot+"/include/inc_footer.html", view.Webroot+"/include/inc_header.html", view.Webroot+"/include/inc_menu.html")
	vars := mux.Vars(r)
	serie := vars["serie"]
	nome := vars["nome"]
	bixo, err := ibc.GerarPagina("ibcs/"+serie+"/"+nome+".json")
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	t.Execute(w, &bixo)
}