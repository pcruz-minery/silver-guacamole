package user

import (
	"context"
	"errors"

	"github.com/mineryreport/flexicar/internal/v1/repository"
	"github.com/mineryreport/flexicar/pkg/model"
)

var ErrNotFound = errors.New("not found")

type userRepository interface {
  Get(ctx context.Context, id string) (*model.User, error)
}

type Controller struct {
  repo userRepository
}

func New(repo userRepository) *Controller {
  return &Controller{
    repo: repo,
  }
}

func (c *Controller) Get(ctx context.Context, id string) (*model.User, error) {
  res, err := c.repo.Get(ctx, id)
  if err != nil && errors.Is(err, repository.ErrNotFound) {
    return nil, ErrNotFound
  }
  return res, err
}
