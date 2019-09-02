package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/trungnguyengtbt/shopping/app/domain/entities"
	"log"
	"os"
)

func Init() *gorm.DB {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", url)
	//defer db.Close()
	if err != nil {
		log.Println(url)
		log.Println(err)
		log.Panic("Error mysql connection")
	}

	//Tables(db, &entities.User{})
	return db
}

func Tables(db *gorm.DB, tables ...interface{}) {
	if db == nil {
		log.Panic("you have to Init database first")
	}
	db.AutoMigrate(tables...)
}
