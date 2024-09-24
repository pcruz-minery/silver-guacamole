package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"github.com/mineryreport/flexicar/pkg/discovery"
)

type Registry struct {
  client *consul.Client
}

func NewRegistry(addr string) (*Registry, error) {
  config := consul.DefaultConfig()
  config.Address = addr
  client, err := consul.NewClient(config)
  if err != nil {
    return nil, err
  }
  return &Registry{client: client}, nil
}

func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
  parts := strings.Split(hostPort, ":")
  if len(parts) != 2 {
    return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8080")
  }
  port, err := strconv.Atoi(parts[1])
  if err != nil {
    return err
      }
  return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
    Address: parts[0],
    ID: instanceID,
    Name: serviceName,
    Port: port,
Check: &consul.AgentServiceCheck{CheckID: instanceID, TTL: "5s"},
  })
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
  services, _, err := r.client.Health().Service(serviceName, "", true, nil)
  if err != nil {
    return nil, err
  } else if len(services) == 0 {
    return nil, discovery.ErrNotFound
  } 

  var addrs []string
  for _, service := range services {
    addrs = append(addrs, fmt.Sprintf("%s:%d", service.Service.Port))
  }

  return addrs, nil
}

func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
  return r.client.Agent().PassTTL(instanceID, "")
}

func (r *Registry) ReportHealthyState(instanceID string,_ string) error {
return r.client.Agent().PassTTL(instanceID, "")
}
