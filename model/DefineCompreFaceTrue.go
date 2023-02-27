package model

import "sicFaceBridge/database"

func DefineCompreFaceTrue(fotoId uint) {
	sql := "UPDATE fotos set compreface = 1 where tatuagem_id = ?"
	db := database.Connect()
	query_erro, err := db.Prepare(sql)
	if !TrataErro(err) {
		query_erro.Exec(fotoId)
	}
}
