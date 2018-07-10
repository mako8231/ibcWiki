package ibc

import (
	"strings"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ibcWiki/view"
	"io/ioutil"
)

type lista struct{
	Ibcs []string 
	Serie string 
}

func ListaDeIbcsHandler(w http.ResponseWriter, r *http.Request){
	var pagina lista 
	files := view.GenerateFiles(view.Webroot+"ibc-list.html", view.Webroot+"include")
	t := view.LoadTemplate(files...)
	index := mux.Vars(r)
	ibcs := PegarLista(index["numero"])
	pagina = lista{Ibcs: ibcs, Serie: index["numero"]}
	t.Execute(w, &pagina)
}

func PegarLista(index string)[]string{
	
	list, err := ioutil.ReadDir("ibcs/"+index)
	if err != nil{
		fmt.Println(err.Error())
		return nil
	}

	v := make([]string, 0)
	for _, names := range list{
		nome := strings.Replace(names.Name(), ".json", "", -1)
		v = append(v, nome)
	}

	return v 
}