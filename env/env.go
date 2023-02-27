package env

import (
	"sicFaceBridge/controllers"

	"github.com/joho/godotenv"
)

func CarregaVariaveisDeAmbiente() {
	err := godotenv.Load()
	controllers.TrataErro(err)
}
