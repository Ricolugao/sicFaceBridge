package model

import (
	"time"
)

type ListaInfratores struct {
	Id         uint      `json:"id"`
	Nome       string    `json:"nome"`
	Acessos    int       `json:"acessos"`
	Vulgo      string    `json:"vulgo"`
	Delitos    string    `json:"delitos"`
	Nascimento time.Time `json:"nascimento"`
	Idade      int       `json:"idade"`
}

type Foto struct {
	Id         uint
	InfratorId int
	Arquivo    string
}

type Infrator struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome"`
	Rg    string `json:"rg"`
	Cpf   string `json:"cpf"`
	Fotos []Foto `json:"fotos"`
}
