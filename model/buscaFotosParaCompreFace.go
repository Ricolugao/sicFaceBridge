package model

import (
	"fmt"
	"runtime"
	"sicFaceBridge/database"
)

func BuscaFotosParaCompreFace(quantidade uint) []Foto {
	db := database.Connect()
	defer db.Close()

	var fotos []Foto

	sql := fmt.Sprintf("select tatuagem_id, infrator_id from fotos where tipo_foto <> 'perfil' and not face limit %d", quantidade)
	// and tipo_foto <> 'perfil'

	rows, err := db.Query(sql)
	if erro := TrataErro(err); !erro {
		defer rows.Close()

		for rows.Next() {
			var foto Foto
			if err := rows.Scan(&foto.Id, &foto.InfratorId); err != nil {
				TrataErro(err)
			}
			fotos = append(fotos, foto)
		}

	}

	return fotos
}

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
