package db

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/trungnguyengtbt/core/common/cons"
)

func Init(conf DbConfiguration) *gorm.DB {
	username := conf.DbUsername
	password := conf.DbPassword
	dbHost := conf.DbHost
	dbPort := conf.DbPort
	dbName := conf.DbName
	dbDriverName:=conf.DbDriverName//default mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, dbHost, dbPort, dbName)

	db, err := gorm.Open(dbDriverName, url)
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
