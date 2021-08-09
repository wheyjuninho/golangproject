package model

import (
	"time"
)

type Sensor struct {
	//gorm.Model
	SensorValue int        `json:"sensorValue"`
	Id1         int        `json:"id1"`
	Id2         string     `json:"id2"`
	Timestamp   *time.Time `json:"timestamp"`
}
