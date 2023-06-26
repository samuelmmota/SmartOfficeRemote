package config

import (
	"go_webserver/entity"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// usar dentro dos serviços
var Db *gorm.DB

func ConnectDB() {
	/*err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}*/

	var err error

	dbUser := os.Getenv("DB_USERNAME")
	// If a required environment variable is not set, handle the error appropriately
	if dbUser == "" {
		log.Fatal("DB_USERNAME environment variable is not set")
	}
	dbPass := os.Getenv("DB_PASSWORD")
	// If a required environment variable is not set, handle the error appropriately
	if dbPass == "" {
		log.Fatal("DB_PASSWORD environment variable is not set")
	}
	dbHost := os.Getenv("DB_HOST")
	// If a required environment variable is not set, handle the error appropriately
	if dbHost == "" {
		log.Fatal("DB_HOST environment variable is not set")
	}
	dbPort := os.Getenv("DB_PORT")
	// If a required environment variable is not set, handle the error appropriately
	if dbPort == "" {
		log.Fatal("DB_PORT environment variable is not set")
	}
	dbDatabase := os.Getenv("DB_DATABASE")
	// If a required environment variable is not set, handle the error appropriately
	if dbDatabase == "" {
		log.Fatal("DB_DATABASE environment variable is not set")
	}

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!" + err.Error())
	}

	err = Db.AutoMigrate(&entity.Mosquitto{}, &entity.SensorData{})
	if err != nil {
		panic("Failed to migrate database!")
	}

}

// fechar a ligação a base de dados
func CloseDb() {
	db, err := Db.DB()
	if err != nil {
		panic("Failed to close database!")
	}
	db.Close()
}

// Check if the database connection is established
func IsDbConnected() bool {
	return Db != nil
}
