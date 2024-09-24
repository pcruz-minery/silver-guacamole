package database

import (
	"log"

	"gorm.io/gorm"
  "gorm.io/driver/postgres"
)

func DbConn() *gorm.DB {
  db, err := gorm.Open(
    postgres.Open("host=localhost user=postgres password=postgres dbname=flexicar port=5432 sslmode=disable"),
    &gorm.Config{},
  )

  if err != nil {
    log.Fatalf("There was an error connecting to the database: %v", err)
  }
  return db
}

