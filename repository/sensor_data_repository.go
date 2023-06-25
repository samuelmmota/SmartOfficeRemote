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

func InsertSensorData(user entity.SensorData) entity.SensorData {
	config.Db.Save(&user)
	return user
}
