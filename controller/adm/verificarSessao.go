package adm

import (
	"net/http"
	"ibcWiki/view"
	"log"
)
//Verifica se o usuário está logado ou não
func VerificarSessao(r *http.Request) bool{
	nome := pegarUsername(r)
	if nome == ""{
		return false 
	}

	if nome != usr{
		return false  
	}

	return true 
}

func pegarUsername(r *http.Request) string{
	cookie, err := r.Cookie(sessao)
	if err != nil{
		log.Println(err.Error())
		return ""
	}

	valor := make(map[string]string)

	err = view.CookieHandler.Decode(sessao, cookie.Value, &valor)
	if err != nil{
		log.Println(err.Error())
		return ""
	}

	return valor[nome]

}
