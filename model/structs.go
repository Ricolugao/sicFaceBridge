package model

type Foto struct {
	Id         uint
	InfratorId int
	Arquivo    string
}

type Retorno struct {
	ImageId string
	Subject string
	Message string
	Code    int
}
