package service

import (
	"go_webserver/entity"
	"go_webserver/repository"
)

func GetAllSensorData() []entity.SensorData {
	return repository.GetAllSensorData()
}

func InsertSensorData(sensorData entity.SensorData) entity.SensorData {
	sensorData = repository.InsertSensorData(sensorData)
	return sensorData
}

/*
func DeleteSensorData(userID uint64) error {
	if err := repository.DeleteUser(userID); err == nil {
		return nil
	}
	return errors.New("User do not exists")
}
*/
