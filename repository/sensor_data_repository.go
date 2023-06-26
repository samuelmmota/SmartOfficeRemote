package repository

import (
	"go_webserver/config"
	"go_webserver/entity"
)

func GetAllSensorData() []entity.SensorData {
	var sensorDataList []entity.SensorData
	config.Db.Preload("SensorData").Find(&sensorDataList)

	return sensorDataList
}

func InsertSensorData(sensorData entity.SensorData) entity.SensorData {
	config.Db.Save(&sensorData)
	return sensorData
}
