package view

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
)

var (
	Route = mux.NewRouter().StrictSlash(true)
	Webroot = "htdocs/"
	ADMRoot = "admDocs/"
	CookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
	Funcs = template.FuncMap{
		"maior" : func(x int, y int) bool {
			return x > y 
		},

		"menor" : func(x int, y int) bool {
			return x < y 
		},
	
		"min" : func(x int, y int) int {
			return x - y
		},
	
		"plus" : func(x int, y int) int {
			return x + y 
		},
	}
)