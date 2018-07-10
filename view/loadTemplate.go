package view

import (
	"html/template"
)

func LoadTemplate(files ...string) *template.Template{
	t, err := template.ParseFiles(files...)
	if err != nil{
		t,_ = template.New("Error").Parse("Error loading template: "+err.Error())
		return t
	}

	return t 
}