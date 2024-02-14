package Config

import (
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }
}

func ConnectionDatabase() {
    DB_HOST := os.Getenv("DB_HOST")
    DB_PORT := os.Getenv("DB_PORT")
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_DATABASE := os.Getenv("DB_DATABASE")

    dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?parseTime=true"
    database, err := gorm.Open(mysql.Open(dsn))
    if err != nil {
        panic(err)
    }

    DB = database
}
