package memory

import (
	"context"
	"sync"

	"github.com/mineryreport/flexicar/internal/v1/repository"
	"github.com/mineryreport/flexicar/pkg/model"
)

type Repository struct {
  sync.RWMutex
  data map[string]*model.User
}

func New() *Repository {
  return &Repository{
    data: map[string]*model.User{
      "1": {ID: "1", Name: "Alice"},
    },
  }
}

func (r *Repository) Get(_ context.Context, id string) (*model.User, error) {
  r.RLock()
  defer r.RUnlock()
  m, ok := r.data[id]

  if !ok {
    return nil, repository.ErrNotFound
  }

  return m, nil
}

func (r *Repository) Put(_ context.Context, id string, user *model.User) error {
  r.Lock()
  defer r.Unlock()
  r.data[id] = user
  return nil
}
