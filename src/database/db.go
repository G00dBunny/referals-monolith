package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(){
	_, err := gorm.Open(mysql.Open("myuser:mypassword@tcp(host.docker.internal:3306)/mydatabase"), &gorm.Config{})

	if err != nil {
		panic( "Could not connect with db")
	}
}