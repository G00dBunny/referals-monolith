package database

import (
	"referals/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var err error
	DB, err = gorm.Open(mysql.Open("myuser:mypassword@tcp(host.docker.internal:3306)/mydatabase"), &gorm.Config{})

	if err != nil {
		panic( "Could not connect with db")
	}
}

func AutoMigrate(){
	DB.AutoMigrate(models.User{})
}