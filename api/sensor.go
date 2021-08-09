package api

import (
	"junlow/db"
	"junlow/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetSensorData(c echo.Context) error {
	db := db.DbManager()
	sensorList := []model.Sensor{}

	id1 := c.QueryParam("Id1")
	id2 := c.QueryParam("Id2")
	if id1 != "" && id2 != "" {

		db = db.Where("Id1 = ? and Id2 = ?", id1, id2)
	}

	db.Find(&sensorList)

	return c.JSON(http.StatusOK, sensorList)
}

func GetFilteredSensorData(c echo.Context) error {
	Id1 := c.QueryParam("Id1")
	Id2 := c.QueryParam("Id2")

	db := db.DbManager()
	sensorList := []model.Sensor{}
	db.Where("Id1 = ? and Id2 = ?", Id1, Id2).Find(&sensorList)

	return c.JSON(http.StatusOK, sensorList)
}
