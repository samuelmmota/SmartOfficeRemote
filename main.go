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
	/*if config.IsDbConnected() {
		// Database connection is established, proceed with importing CSV
		utils.ImportCSV()
	} else {
		// Database connection failed, handle the error accordingly
		log.Fatal("Failed to connect to the database")
	}*/

	// Close the database connection when done
	config.CloseDb()
}
