package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct{
	Port string `json:"port"`
	Admin Admin `json:"admin"`
}

type Admin struct{
	Password string `json:"password"`
	Username string `json:"username"`
}

var (
	Server *Config
)

func ReadConfig(){
	//Carregando o arquivo
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil{
		log.Fatal("Falha ao carregar arquivo JSON: "+err.Error())
	}

	//Pegando todos os dados
	err = json.Unmarshal(b, &Server)
	if err != nil{
		log.Fatal("Falha ao ler arquivo de configuração: "+err.Error())
	}

}