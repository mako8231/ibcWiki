package ibc

import (
	"strings"
	"net/http"
	"github.com/gorilla/mux"
	"errors"
	"io/ioutil"
	"ibcWiki/model"
)

func EditarIBC(r *http.Request) (err error){
	//Pegar as variáveis passadas pela url
	var msg string 
	var legendasV = []string{}
	var galeriaV = []string{}
	vars := mux.Vars(r)
	serie := vars["serie"]
	nome := vars["nome"]

	//Verificar se o arquivo existe
	existe := arquivoExiste(serie, nome)

	/*Caso contrário a função irá parar no mesmo instante, pois é
	impossível editar um arquivo inexistente 
	*/
	if !existe{
		return errors.New("arquivo inexistente")
	}

	//Agora, pegaremos as variáveis do formulário
	apelido := r.FormValue("txtApelido")
	classe := r.FormValue("txtClasse")
	descricao := r.FormValue("txtDescricao")
	imagem := r.FormValue("txtImagem")
	legendaImagem := r.FormValue("txtLegendaImagem")	
	galeriaImagens := r.FormValue("txtGaleriaImagens")
	galeriaLegendas := r.FormValue("txtGaleriaLegendas")

	//Verificar se os campos obrigatórios estão preenchidos
	if apelido==""{msg+="apelido não informado \n"}
	if classe==""{msg+="classe não informada \n"}
	if descricao==""{msg+="descrição não informada \n"}
	if imagem==""{msg+="imagem não informada"}
	if legendaImagem==""{msg+="legenda da imagem não informada \n"}

	//Separar as legendas em uma array
	if galeriaLegendas != ""{
		legendasV = strings.Split(galeriaLegendas, "|")
	}

	if galeriaImagens != ""{
		galeriaV = strings.Fields(galeriaImagens)
	}

	//Se tudo está ok, vamos começar com a alteração
	if msg == ""{
		//a partir da classe de objeto informada: pegar o valor da cor
		cor, nomeClasse := pegarCor(classe)
		ibcObj := model.Ibc{Dados: []string{nome, apelido, descricao, imagem, legendaImagem, nomeClasse, cor}, Legendas: legendasV}
		//Tem galeria?
		if galeriaImagens != ""{
			ibcObj = inserirGaleria(galeriaV, ibcObj)
		}
		err = SalvarJson(ibcObj, serie, nome)
		if err != nil{
			return err
		}	
	} else {
		return errors.New(msg)
	}


	return nil 
}

func arquivoExiste(serie, nome string) bool{
	_, err := ioutil.ReadFile("ibcs/"+serie+"/"+nome+".json")
	if err != nil{
		return false 
	}

	return true
}