package config

import (
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlPort := os.Getenv("MYSQL_PORT")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDbname := os.Getenv("MYSQL_DBNAME")

    config := map[string]string{
        "MYSQL_USER":     mysqlUser,
        "MYSQL_PASSWORD": mysqlPassword,
        "MYSQL_PORT":     mysqlPort,
        "MYSQL_HOST":     mysqlHost,
        "MYSQL_DBNAME":   mysqlDbname,
    }

    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
        config["MYSQL_USER"],
        config["MYSQL_PASSWORD"],
        config["MYSQL_HOST"],
        config["MYSQL_PORT"],
        config["MYSQL_DBNAME"])

    db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return db
}
