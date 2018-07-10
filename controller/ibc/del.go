package ibc

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func DeleteIBC(r *http.Request) error {
	//url: /mod/painel/{serie}/{nome}/deletar
	vars := mux.Vars(r)
	err := deletarArquivo(vars["serie"], vars["nome"])
	if err != nil{
		return err 
	}

	return nil 
}

func deletarArquivo(pasta, nome string)(error){
	err := os.Remove("ibcs/"+pasta+"/"+nome+".json")
	if err != nil{
		return err 
	}

	return nil 
}

