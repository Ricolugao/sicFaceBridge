package database

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Connect() sql.DB {
	DbUser := os.Getenv("DbUser")
	DbPassword := os.Getenv("DbPassword")
	DbHost := os.Getenv("DbHost")
	DbPort := os.Getenv("DbPort")
	DbDatabase := os.Getenv("DbDatabase")
	connection, err := sql.Open("mysql", DbUser+":"+DbPassword+"@tcp("+DbHost+":"+DbPort+")/"+DbDatabase+"?parseTime=true")
	// connection, err := sql.Open("mysql", "sic:123456@tcp(sic-mysql:3306)/sic")

	if err != nil {
		_, fileName, linha, _ := runtime.Caller(0)
		erro := fmt.Sprintf("Arquivo: %s, Linha: %d, Erro: %s", fileName, linha, err.Error())
		arquivo, err := os.OpenFile("LogDeErros.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Falha no arquivo LogDeErros.txt:", err)
		}
		log := fmt.Sprintf(time.Now().Format("02/01/2006 15:04:05")+"\nErro: %s  |  %s  |  %s\n\n", erro, err.Error())
		arquivo.WriteString(log)
		arquivo.Close()
	}
	if err != nil {
		fmt.Println("erro ao criar handle do banco de dados:", err.Error())
	}

	// VefificaConexao(*connection)

	connection.SetConnMaxLifetime(time.Second * 5)
	connection.SetMaxOpenConns(100)
	connection.SetMaxIdleConns(100)
	return *connection
}

func VefificaConexao(db sql.DB) bool {
	resposta := false
	err := db.Ping()
	if err != nil {
		fmt.Println("Erro de conexao: ", err.Error())
		resposta = true
	}
	return resposta
}
