package main

import (
	"junlow/db"
	route "junlow/routes"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
