package initializers

import (
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
    dsn := "root:iyacorp123@tcp(127.0.0.1:3306)/mnc?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Could not connect to database: ", err)
    }
}
