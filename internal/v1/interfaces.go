package v1

import "github.com/mineryreport/flexicar/pkg/model"

type RepositoryInterface interface {

  Create(model.User) (model.User, error)
  Get(id string) (model.User, error)
  Update(model.User) error
  Delete(id string) error
  GetByEmail(email string) (model.User, error)
}

type UseCaseInterface interface {
  Create(model.User) (model.User, error)

  Get(id string) (model.User, error)

  Update(model.User) error

  Delete(id string) error


}
