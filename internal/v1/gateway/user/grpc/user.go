package grpc

import (
	"context"

	"github.com/mineryreport/flexicar/gen"
	"github.com/mineryreport/flexicar/internal/v1/grpcutil"
	"github.com/mineryreport/flexicar/pkg/discovery"
	"github.com/mineryreport/flexicar/pkg/model"
)

type Gateway struct {
registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
  return &Gateway{registry }
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.User, error) {
  conn, err := grpcutil.ServiceConnection(ctx, "user", g.registry)
  if err != nil {
    return nil, err
  }

  defer conn.Close()

  client := gen.NewUserServiceClient(conn)
  
  resp, err := client.Read(ctx, &gen.SingleUserRequest{Id: id})
  if err != nil {
    return nil, err
      }
  return model.UserFromProto(resp) , nil
}

