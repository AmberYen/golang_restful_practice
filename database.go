package main

import (
  "log"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
  // db connection
  db, err := gorm.Open("mysql", "rytass:rytass2O15@/go_articles?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    log.Printf("%s", err)
  }
  
  db.LogMode(true)

  return db
}
