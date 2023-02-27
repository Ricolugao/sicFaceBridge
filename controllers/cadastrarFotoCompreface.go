package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"sicFaceBridge/model"
	"strconv"
	"strings"
)

func CadastraFotoCompreFace(msg []byte) model.Retorno {

	var instrucoes model.Foto
	err := json.Unmarshal(msg, &instrucoes)
	TrataErro(err)
	foto := model.BuscaFotosDeInfratores(instrucoes.Id)
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
	TrataErro(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", "ca4b55ff-c571-4e13-874b-4e44303e16af")

	res, err := client.Do(req)
	TrataErro(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	TrataErro(err)
	var retorno model.Retorno
	err = json.Unmarshal(body, &retorno)
	TrataErro(err)
	if retorno.Code == 0 {
		model.DefineCompreFaceTrue(foto.Id)
	}
	return retorno
}
