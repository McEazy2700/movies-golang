package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb()(*gorm.DB, error){
  dsn := "host=localhost user=Eazy password=gorm dbname=go-moviesdb port=5432"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  db.AutoMigrate()
  return db, nil
}
