package database

import (
	"first_go/src/model"
	"fmt"

	"gorm.io/gorm"
)

var DATABASE *gorm.DB

type dbConfig struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
}

func Build() *dbConfig {
	config := dbConfig{
		Host:     "localhost",
		Port:     8000,
		User:     "root",
		Password: "",
		Name:     "isomorphic_tb",
	}
	return &config
}

func Url(config *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}

func DoMigration() {
	DATABASE.AutoMigrate(&model.User{})
	DATABASE.AutoMigrate(&model.MigrationTeacher{})
}
