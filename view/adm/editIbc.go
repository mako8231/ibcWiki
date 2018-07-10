package adm

import (
	"github.com/gorilla/mux"
	"net/http"
	"ibcWiki/view"
	"ibcWiki/controller/adm"
	"ibcWiki/controller/ibc"
	"io/ioutil"
	"html/template"
	"ibcWiki/model"
)

type editPage struct {
	Serie string 
	Bixo model.Ibc
}

func EditIBCFormHandler(w http.ResponseWriter, r *http.Request){
	var pagina editPage
	logado := adm.VerificarSessao(r)
	if logado{
		//Abrindo o arquivo principal
		b, err := ioutil.ReadFile(view.ADMRoot+"ibc-edit.html")
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotFound)
			return 
		}


		//Pegando o valor das variáveis passadas pela URL 
		vars := mux.Vars(r)
		serie := vars["serie"]
		nome := vars["nome"]
		//Pegando o ibc salvo
		bixo, err := ibc.GerarPagina("ibcs/"+serie+"/"+nome+".json")
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotFound)
			return 
		}
		//Lendo arquivos
		t,_ := template.Must(template.New("ibc").Funcs(view.Funcs).Parse(string(b))).ParseFiles(view.ADMRoot+"/include/footer.html", view.ADMRoot+"/include/header.html", view.ADMRoot+"/include/menu.html")
		pagina = editPage{Serie: serie, Bixo: bixo}
		t.Execute(w, &pagina)
	} else {
		http.Redirect(w, r, "/mod/", 302)
	}
}

func AlterarIBCHandler(w http.ResponseWriter, r *http.Request){
	logado := adm.VerificarSessao(r)
	if logado{
		err := ibc.EditarIBC(r)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}

		http.Redirect(w, r, "/mod/", 302)
	}
}