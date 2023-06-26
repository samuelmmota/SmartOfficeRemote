package utils

import (
	"encoding/csv"
	"fmt"
	"go_webserver/entity"
	"go_webserver/repository"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func ImportCSV() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %s", err.Error())
	}

	// Define the relative path to the CSV file
	relativePath := "exports/csv_templates/dht11_data.csv"

	// Construct the absolute path to the CSV file
	filePath := filepath.Join(wd, relativePath)

	// Open the CSV file
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Failed to open CSV file:", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV records:", err)
	}

	// Iterate over each record and create SensorData objects
	var sensorDataList []entity.SensorData
	for _, record := range records {
		// Parse the values from the record
		id, _ := strconv.ParseUint(record[0], 10, 64)
		temperature, _ := strconv.ParseFloat(record[1], 64)
		humidity, _ := strconv.ParseFloat(record[2], 64)
		name := record[3]
		dataType := record[4]
		location := record[5]
		timestamp, _ := time.Parse("2006-01-02 15:04:05", record[6])

		// Create a new SensorData object
		sensorData := entity.SensorData{
			SensorID:    id,
			Temperature: temperature,
			Humidity:    humidity,
			Name:        name,
			Type:        dataType,
			Location:    location,
			Timestamp:   timestamp,
		}

		// Add the SensorData object to the list
		sensorDataList = append(sensorDataList, sensorData)
	}

	// Print the imported sensor data
	for _, data := range sensorDataList {
		fmt.Printf("ID: %d, Temperature: %.2f, Humidity: %.2f, Name: %s, Type: %s, Location: %s, Timestamp: %s\n",
			data.ID, data.Temperature, data.Humidity, data.Name, data.Type, data.Location, data.Timestamp)
		repository.InsertSensorData(data)
	}
}
