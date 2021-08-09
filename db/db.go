package db

import (
	"fmt"
	"junlow/config"
	"junlow/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	var configuration = config.GetConfig()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connectionString)

	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}

	db.AutoMigrate(model.Sensor{})
}

func DbManager() *gorm.DB {
	return db
}
