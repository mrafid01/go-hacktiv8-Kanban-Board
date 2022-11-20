package config

import (
	"fmt"
	"log"
	"project3/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(username, password, host, port, dbName string) *gorm.DB {

	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}
	if err != nil {
		log.Fatal("DB Konek Eror")
	}
	fmt.Println("DB Berhasil Konek")
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Task{})
	db.AutoMigrate(&entity.Category{})
	return db
}
