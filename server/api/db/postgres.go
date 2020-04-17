package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nickmurr/golang-api/config"
)

var (
	DBDriver = "postgres"
	DBURL    = "host=localhost port=5432 user=postgres dbname=postgres password=docker sslmode=disable"
	DB       *gorm.DB
)

func Load() *gorm.DB {
	db, err := gorm.Open(DBDriver, DBURL)
	if err != nil {
		config.Log.Fatal(err)
		return nil
	}
	config.Log.Info("Success connection to POSTGRES")
	DB = db
	return db
}
