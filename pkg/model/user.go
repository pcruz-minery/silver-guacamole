package model

import "gorm.io/gorm"

type User struct {
  gorm.Model

  ID   string `gorm:"primaryKey";default:"uuid_generate_v4()"`
  Name string 
  Email string `gorm:"unique;not nul"`
}
