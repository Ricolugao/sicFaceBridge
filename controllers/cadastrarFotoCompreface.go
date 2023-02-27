package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CadastraFotoCompreFace(msg []byte) {

	// arquivoBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	mensagem := strings.Split(string(msg), ";")

	foto := mensagem[1]
	infratorId := mensagem[0]

	// fileBase64 := base64.StdEncoding.EncodeToString([]byte(foto))

	url := "http://192.168.144.17:8000/api/v1/recognition/faces?subject=" + infratorId
	method := "POST"
	payload := strings.NewReader(`{` + "  " + `	"file":"` + foto + `"` + "  " + `  }`)

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

	fmt.Println("retorno: ", []byte(body))

}
