package connection

import (
	"book-crud/config"
	"book-crud/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connect() {
	dbConfig := config.InitConfig()
	dsn := fmt.
		Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIP, dbConfig.DbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

func migrate() {
	db.Migrator().AutoMigrate(&model.Book{})
}

func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	migrate()
	return db
}
