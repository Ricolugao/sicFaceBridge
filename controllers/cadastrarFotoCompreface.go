package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sicFaceBridge/database"
	"sicFaceBridge/model"
	"strconv"
	"strings"
)

func CadastraFotoCompreFace(msg []byte) {

	// arquivoBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	fmt.Println(string(msg))
	var instrucoes model.Foto
	err := json.Unmarshal(msg, &instrucoes)
	TrataErro(err)
	foto := BuscaFotosDeInfratores(int(instrucoes.Id))[0]
	foto.InfratorId = instrucoes.InfratorId
	foto.Id = instrucoes.Id
	// mensagem := strings.Split(string(msg), ";")

	// foto := mensagem[1]
	infratorId := strconv.Itoa(foto.InfratorId)

	fileBase64 := base64.StdEncoding.EncodeToString([]byte(foto.Arquivo))

	url := "http://192.168.144.1:8000/api/v1/recognition/faces?subject=" + infratorId
	method := "POST"
	payload := strings.NewReader(`{` + "  " + `	"file":"` + fileBase64 + `"` + "  " + `  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", "ca4b55ff-c571-4e13-874b-4e44303e16af")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("retorno: ", string(body))

}

func BuscaFotosDeInfratores(FotoId int) []model.Foto {
	db := database.Connect()
	defer db.Close()

	var fotos []model.Foto

	sql := fmt.Sprintf("select arquivo from fotos where tatuagem_id = %d", FotoId)
	// and tipo_foto <> 'perfil'
	rows, err := db.Query(sql)
	if erro := TrataErro(err); !erro {
		defer rows.Close()

		for rows.Next() {
			var foto model.Foto
			if err := rows.Scan(&foto.Arquivo); err != nil {
				TrataErro(err)
			}
			fotos = append(fotos, foto)
		}

	}

	return fotos
}
