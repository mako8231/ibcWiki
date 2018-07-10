package ibc

import (
	"ibcWiki/model"
	"io/ioutil"
	"encoding/json"
)

func GerarPagina(caminho string) (model.Ibc, error){
	var ibc model.Ibc
	b, err := ioutil.ReadFile(caminho)
	if err != nil{
		return model.Ibc{}, err 
	}

	err = json.Unmarshal(b, &ibc)
	if err != nil{
		return model.Ibc{}, err 
	}

	return ibc, nil 
}