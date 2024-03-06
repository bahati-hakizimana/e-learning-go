package database

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get database DNS from environment variable
    dsn := os.Getenv("DSN")

    // Connect to the database
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Could not connect to database: " + err.Error())
    }else{
		log.Println("Successful Connected")
	}

    DB = database
}
