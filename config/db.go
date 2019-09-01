package config 

import (
  "os"
  "fmt"
  "github.com/jinzhu/gorm"
  "searchCurator/models"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
  dbHost, _ := os.LookupEnv("POSTGRES_HOST")
  username, _ := os.LookupEnv("POSTGRES_USER")
  dbName, _ := os.LookupEnv("POSTGRES_DB")
  password, _ := os.LookupEnv("POSTGRES_PASSWORD")
  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
  db, err := gorm.Open("postgres", dbUri)
  if err != nil {
    panic(err)
  }
  db.LogMode(true)
  db.AutoMigrate(&models.Listing{})
  db.AutoMigrate(&models.Item{})
  return db
}

func CloseConnection(db *gorm.DB) {
  defer db.Close()
}
