package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL1")), &gorm.Config{})

	if err != nil {
		panic("Could not connect to DATABASE")
	}
}
