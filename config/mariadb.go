package config

import (
	"fmt"

	"github.com/maxzycon/SIB-Golang-Assigment-2/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMariaDb() *gorm.DB {
	username := "root"
	pass := "oskar101"
	dbName := "mbkm_assigment_1"

	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		username, pass,
		"localhost", dbName,
	)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
	return db
}
