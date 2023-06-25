package entity

import "time"

type SensorData struct {
	ID          uint64 `gorm:"primary_key:auto_increment"`
	SensorID    uint64
	Temperature float64
	Humidity    float64
	Name        string `gorm:"type:varchar(255)"`
	Type        string `gorm:"type:varchar(255)"`
	Location    string `gorm:"type:varchar(255)"`
	Timestamp   time.Time
}
