package adm

import (
	"net/http"
	"ibcWiki/view"
)

//Após a validação dos dados, o servidor cria uma chave criptografada como valor passado para o cookie
func CriarSessao(username string, w http.ResponseWriter) error {
	valor := map[string]string{
		nome : username,
	}

	codificado, err := view.CookieHandler.Encode(sessao, valor)
	if err != nil{
		return err 
	} 

	cookie := &http.Cookie{
		Name: sessao,
		Value: codificado,
		Path: caminho,
	}	

	http.SetCookie(w, cookie)
	return nil
}