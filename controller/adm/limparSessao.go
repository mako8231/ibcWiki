package adm

import (
	"net/http"
)

//Basicamente, o Logout 
func LimparSessao(w http.ResponseWriter){
	cookie := &http.Cookie{
		Name: sessao,
		Value: "",
		Path: caminho,
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}