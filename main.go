package main

import (
	"first_go/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	err = godotenv.Load()
	if err != nil {
		panic("Failed To Load .env File")
	}

	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	database.DATABASE, err = gorm.Open(mysql.Open(dbUrl))

	if err != nil {
		panic("Failed To Connect Database")
	}
	database.DoMigration()

	fmt.Println("woke")
}
