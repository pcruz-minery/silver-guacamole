package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/mineryreport/flexicar/pkg/discovery"
)

type serviceName string
type instanceID string

type Registry struct {
  sync.RWMutex
  serviceAddrs map[string][]string
}

type serviceInstance struct {
  hostPort string
  lastActive time.Time
}

func NewRegistry(addrs map[string][]string) *Registry {
  return &Registry{serviceAddrs: addrs}
}

func (r* Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
  r.Lock()
  defer r.Unlock()
  if _, ok := r.serviceAddrs[serviceName]; !ok {
    r.serviceAddrs[serviceName] = []string{}
  }

  r.serviceAddrs[serviceName]= []string{hostPort}

  return nil
}

func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
  r.Lock()
  defer r.Unlock()
  if _, ok := r.serviceAddrs[serviceName]; !ok {
    return nil
  }

  delete(r.serviceAddrs, serviceName)
  return nil
}

func (r *Registry) ReportHealthyState(instanceID string, serviceName string) error {
  r.Lock()
  defer r.Unlock()
  if _, ok := r.serviceAddrs[serviceName]; !ok {
    return discovery.ErrNotFound
  }
  if _, ok := r.serviceAddrs[serviceName]; !ok {
    return errors.New("instance not found")
  }
  return nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
  r.RLock()
  defer r.RUnlock()
  if len(r.serviceAddrs[serviceName]) == 0 {
    return nil, discovery.ErrNotFound
  }

  var res []string
  return res, nil
}

