package db

import (
	"fmt"
	"gin-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := "localhost"
	port := "3306"
	database := "go_db"
	username := "root"
	password := ""
	charset := "utf8"
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to DB connection ,err" + err.Error())
	}
	autoMigrateTable()
}

func GetDB() *gorm.DB {
	return DB
}

func autoMigrateTable() {
	DB.AutoMigrate(&models.User{}, &models.Author{}, &models.Blog{})
	fmt.Println("自动迁移数据表")
}
