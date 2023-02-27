package controllers

import (
	"fmt"
	"runtime"
)

func TrataErro(err error) bool {
	retorno := false
	_, arquivo, linha, _ := runtime.Caller(1)
	erro := fmt.Sprintf("Arquivo: %s, Linha: %d", arquivo, linha)
	if err != nil {
		fmt.Println(erro, err.Error())
		retorno = true
	}
	return retorno
}
