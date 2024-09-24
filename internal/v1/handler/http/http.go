package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mineryreport/flexicar/internal/v1/controller/user"
	"github.com/mineryreport/flexicar/internal/v1/repository"
)

type Handler struct {
  ctrl *user.Controller
}

func New(ctrl *user.Controller) *Handler {
  return &Handler { ctrl }
}

func (h *Handler) GetUser(w http.ResponseWriter, req *http.Request) {
  id := req.FormValue("id")
  

  if id == "" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  ctx := req.Context()
  
  m, err := h.ctrl.Get(ctx, id)

  if err != nil && errors.Is(err, repository.ErrNotFound) {
    w.WriteHeader(http.StatusNotFound)
    return
  } else if err != nil {
    log.Printf("Repository get error: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  if err := json.NewEncoder(w).Encode(m); err != nil {
    log.Printf("Reponse encode error: %v\n", err)
  }
}
