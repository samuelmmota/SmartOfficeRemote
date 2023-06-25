package main

import (
	"go_webserver/config"
	"go_webserver/utils"
)

func main() {

	//configurar a base de dados
	config.ConnectDB()

	//deferrir o fecho da base de dados
	//evitar que feche logo a base de dados
	defer config.CloseDb()

	utils.ImportCSV()
}
