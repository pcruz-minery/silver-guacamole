package http

import (
	"context"
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

  "github.com/mineryreport/flexicar/internal/v1/gateway"
	"github.com/mineryreport/flexicar/pkg/model"
	"github.com/mineryreport/flexicar/pkg/discovery"
)

type Gateway struct {
  registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
  return &Gateway{
    registry ,
  }
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.User, error) { 
  addrs, err := g.registry.ServiceAddresses(ctx, "user")
  if err != nil {
    return nil, err
      }
  url := "http://" + addrs[rand.Intn(len(addrs))] + "/user"
  log.Printf("Calling user service. Request: GET " + url)

  req, err := http.NewRequest(http.MethodGet, url, nil)
  req = req.WithContext(ctx)
  values := req.URL.Query()

  values.Add("id", id)
  req.URL.RawQuery = values.Encode()
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
      }
  defer resp.Body.Close()
  if resp.StatusCode == http.StatusNotFound {
    return nil, gateway.ErrNotFound
      } else if resp.StatusCode/100 != 2 {
    return nil, fmt.Errorf("non-2xx response: %v", resp)
  }
  var v *model.User
  if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
    return nil, err
  }
  return v, nil

}
