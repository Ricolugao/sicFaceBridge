package model

import "sicFaceBridge/database"

func BuscaFotosDeInfratores(FotoId uint) Foto {
	db := database.Connect()
	defer db.Close()

	var foto Foto
	sql := "select arquivo from fotos where tatuagem_id = ?"
	query, err := db.Prepare(sql)
	if erro := TrataErro(err); !erro {
		defer query.Close()
		if err := query.QueryRow(FotoId).Scan(&foto.Arquivo); err != nil {
			TrataErro(err)
		}
	}
	return foto
}
