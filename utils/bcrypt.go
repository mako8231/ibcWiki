package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//Compara se a senha digitada possui uma array de bytes que condiz com o hash
func CompararSenhas(senha, cripto string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cripto), []byte(senha))
	if err != nil{
		return false 
	}

	return true 
}