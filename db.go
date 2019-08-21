package main

import (
  "github.com/jinzhu/gorm"
  "./models"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() *gorm.DB {
  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to open db")
  }

  db.AutoMigrate(&models.Listing{})

  db.Create(&models.Listing{Url:"https://indeed.com", Selector:"body>div"})
  return db
}

func CloseConnection(db *gorm.DB) {
  defer db.Close()
}
