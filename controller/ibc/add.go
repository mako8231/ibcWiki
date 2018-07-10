package ibc

import (
	"io/ioutil"
	"net/http"
	"errors"
	"strings"
	"encoding/json"
	"ibcWiki/model"
	"log"
)

func CadastrarIBC(r *http.Request)(err error){
	var msg string
	var galeriaArray []string
	var galeriaLegendasA []string
	var ibcObj model.Ibc
	log.Println("teste")
	//Pegar todos os dados do formulário
	serie := r.FormValue("txtSerie")
	nome := r.FormValue("txtNome")
	apelido := r.FormValue("txtApelido")
	classe := r.FormValue("txtClasse")
	descricao := r.FormValue("txtDescricao")
	imagem := r.FormValue("txtImagem")
	legendaImagem := r.FormValue("txtLegendaImagem")
	galeria := r.FormValue("txtGaleriaImagens")
	galeriaLegendas := r.FormValue("txtGaleriaLegendas")

	//Controle de fluxo
	if serie == ""{msg += "Série não informada \n"}

	if nome == ""{msg += "Nome não informado \n"}

	if apelido == ""{msg += "Apelido não informado \n"}

	if classe == ""{msg += "Classe de objeto não informado \n"}

	if descricao == ""{msg += "Descrição não informada"}

	if imagem == ""{msg += "Imagem não informada \n"}

	if legendaImagem == ""{msg += "Legenda da imagem não informada \n"}

	if galeria != "" {
		galeriaArray = strings.Fields(galeria)
	}

	if galeriaLegendas != ""{
		galeriaLegendasA = strings.Split(galeriaLegendas, "|")
	}

	//Pegar a cor respectiva do objeto 
	cor, nomeClasse := pegarCor(classe)

	if cor == ""{
		msg += "Classe de objeto inválida \n"
	}

	//Se tudo estiver ok, faça a inserção
	if msg == ""{
		nome = strings.Replace(nome, " ", "", -1)
		ibcObj = model.Ibc{Dados : []string{nome, apelido, descricao, imagem, legendaImagem, nomeClasse, cor}, Legendas: galeriaLegendasA}
		//Se o campo da galeria foi inserido, insira os campos na array
		if galeria != ""{
			ibcObj = inserirGaleria(galeriaArray, ibcObj)
		}
		err := SalvarJson(ibcObj, serie, nome)
			if err != nil{
				log.Println(err.Error())
				return err 
		}
		return nil 
	}

	return errors.New(msg)

}

func SalvarJson(modelo model.Ibc, pasta, nome string) (err error){
	b, err := json.Marshal(modelo)
	if err != nil{
		log.Println(err.Error())
		return err 
	}

	err = ioutil.WriteFile("ibcs/"+pasta+"/"+nome+".json", b, 0666)
	if err != nil{
		log.Println(err.Error())
		return err 
	}

	return nil
}

func pegarCor(nome string) (string, string) {
	switch nome{
	case "1": return "green", "Stefanno"
	case "2": return "green", "Masao"
	case "3": return "yellow", "Arac"
	case "4": return "orange", "Bolu"
	case "5": return "red", "Soul"
	case "6": return "gray", "Darkente"
	case "7": return "blue", "Luan"
	default:
	return "", ""
	}
	return "", ""
}

func inserirGaleria(galeria []string, ibc model.Ibc) model.Ibc{
	for _, imagem := range galeria{
		ibc.Dados = append(ibc.Dados, imagem)
	}

	return ibc
}
