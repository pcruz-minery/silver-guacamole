package db

import (
"gorm.io/gorm"
"github.com/mineryreport/flexicar/pkg/model"
interfaces "github.com/mineryreport/flexicar/internal/v1"
  "github.com/google/uuid"
)

type Repository struct {
  db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepositoryInterface {
  return &Repository{db}
}

func (r *Repository) Create(user model.User) (model.User, error) {
  user.ID = uuid.New().String()
  err := r.db.Create(&user).Error
  return user, err
}


func (r *Repository) Get(id string) (model.User, error) {
  var user model.User
  err := r.db.Where("id = ?", id).First(&user).Error
  return user, err
}

func (r *Repository) Update(user model.User) error {
  var dbUser model.User
  if err := r.db.Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
    return err
  }

  dbUser.Name = user.Name
  err := r.db.Save(dbUser).Error
  return err
}

func (r *Repository) Delete(id string) error {
  err := r.db.Where("id = ?", id).Delete(&model.User{}).Error
  return err
}

func (r *Repository) GetByEmail(email string) (model.User, error) {
  var user model.User
  err := r.db.Where("email = ?", email).First(&user).Error
  return user, err
}
