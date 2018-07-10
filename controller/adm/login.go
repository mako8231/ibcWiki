package adm

import (
	"net/http"
	"ibcWiki/utils"
)

//Faz o controle do login 
func VerificarCredenciais(r *http.Request, w http.ResponseWriter) bool {
	var valido bool 
	username := r.FormValue("txtUsername")
	senhaF := r.FormValue("txtSenha")
	valido = utils.CompararSenhas(senhaF, senha)
	if !valido {
		return valido 
	}
	if username != usr{
		return false  
	}
	err := CriarSessao(username, w)
	if err != nil{
		return false 
	}
	return valido 
}