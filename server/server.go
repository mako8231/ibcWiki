package server

import (
	"net/http"
	"ibcWiki/config"
	"log"
	"ibcWiki/view"
	"ibcWiki/view/ibc"
	"ibcWiki/view/adm"
)

func StartServer(){
	HTTPServerStart()
}

func HTTPServerStart(){
	config.ReadConfig()
	//Criar um sistema de arquivos 
	assetsFS := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	//Painel administrativo
	view.Route.HandleFunc("/mod/", adm.LoginHandler)
	view.Route.HandleFunc("/verificar", adm.VerificarCredenciaisHandler).Methods("POST")
	view.Route.HandleFunc("/mod/painel/", adm.AdmIndexHandler)
	view.Route.HandleFunc("/mod/painel/ibc/{serie}", adm.IBCListaHandler)
	view.Route.HandleFunc("/mod/ibc/add/", adm.CadastroIBCHandler)
	view.Route.HandleFunc("/mod/painel/salvar_ibc", adm.SalvarIBCHandler).Methods("POST")
	view.Route.HandleFunc("/mod/painel/ibc/edit/{serie}/{nome}", adm.EditIBCFormHandler)
	view.Route.HandleFunc("/mod/painel/{serie}/{nome}/alterar", adm.AlterarIBCHandler).Methods("POST")
	view.Route.HandleFunc("/mod/painel/{serie}/{nome}/deletar", adm.DeleteIBCHandler)
	view.Route.HandleFunc("/mod/painel/logout/", adm.LogoutHandler)

	//Handlers:
	view.Route.HandleFunc("/", view.IndexHandler)
	view.Route.HandleFunc("/sobre-nos/", view.SobreNosHandler)

	//ibcs
	view.Route.HandleFunc("/series/{numero}", ibc.ListaDeIbcsHandler)
	view.Route.HandleFunc("/view/{serie}/{nome}", ibc.PaginaIBCHandler)

	//Iniciando servidor
	log.Println("Arquivo de configuração carregado com sucesso")
	log.Println("Seu aplicativo web está rodando na porta "+config.Server.Port)
	http.Handle("/", view.Route)
	log.Fatal(http.ListenAndServe(config.Server.Port, nil))
	}