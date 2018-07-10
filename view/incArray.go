package view

import (
	"fmt"
	"io/ioutil"
)

func GenerateFiles(firstFile, folderName string) []string{
	files, err := ioutil.ReadDir(folderName)
	if err != nil{
		fmt.Printf("%s", "Erro "+err.Error())
		return nil
	}

	v := make([]string, len(files) - len(files))
	v = append(v, firstFile)

	for  _, names := range files{
		v = append(v, folderName+"/"+names.Name())
	}
	return v 

}