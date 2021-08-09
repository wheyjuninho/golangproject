package routes

import (
	"junlow/api"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	//Set Bundle Middleware
	e.Use(echoMw.Recover())
	e.Use(echoMw.Logger())

	e.GET("/", api.GetSensorData)

	return e
}
