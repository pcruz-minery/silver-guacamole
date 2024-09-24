package usecase

import (
	"errors"
	"log"

	interfaces "github.com/mineryreport/flexicar/internal/v1"
	"github.com/mineryreport/flexicar/pkg/model"
	"gorm.io/gorm"
)

type UseCase struct {

  repo interfaces.RepositoryInterface
}

func New(repo interfaces.RepositoryInterface) interfaces.UseCaseInterface {
  return &UseCase{repo}
}

func (uc *UseCase) Create(user model.User) (model.User, error) {
  if _, err := uc.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
    return model.User{}, errors.New("Email already exists")
  }
  return uc.repo.Create(user)
}

func (uc *UseCase) Get(id string) (model.User, error) {
  var user model.User
  var err error

  if user, err = uc.repo.Get(id); err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return model.User{}, errors.New("User not found")
    }
    return model.User{}, err
  }

  return user, nil
}


func (uc *UseCase) Update(updateUser model.User) error {

  var user model.User
  var err error

  if user, err = uc.repo.Get(updateUser.ID); err != nil {
    return err
  }

  if user.Email != updateUser.Email {
    return errors.New("Email cannot be updated")
  }

  err = uc.repo.Update(updateUser)
  if err != nil {
    log.Fatalf("error updating %v", err)
    return err
  }

  return nil
}

func (uc *UseCase) Delete(id string) error {
  var err error
  if _, err = uc.Get(id); err != nil {
    return err
  }

  err = uc.repo.Delete(id)

  if err != nil {
    return err
  }

  return nil
}
