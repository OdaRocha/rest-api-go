package main

import (
	"fmt"

	"api-go-rest/routes"
	"api-go-rest/routes/database"
)

func main() {
	fmt.Println("Iniciando o servidor Rest Go")
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
