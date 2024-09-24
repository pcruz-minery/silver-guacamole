package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	dbConfig "github.com/mineryreport/flexicar/internal/v1/db"
  interfaces "github.com/mineryreport/flexicar/internal/v1"
  repo "github.com/mineryreport/flexicar/internal/v1/repository/db"
  usecase "github.com/mineryreport/flexicar/internal/v1/usecase"
  handler "github.com/mineryreport/flexicar/internal/v1/handler/grpc"

	"github.com/mineryreport/flexicar/pkg/discovery"
	"github.com/mineryreport/flexicar/pkg/discovery/consul"
	"github.com/mineryreport/flexicar/pkg/model"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const serviceName = "user"

func main() {
  var port int
  flag.IntVar(&port, "port", 8081, "API handler port")
  flag.Parse()
  log.Printf("Starting the user service on port %d", port)
  
  db := dbConfig.DbConn()
  migrations(db)

  registry, err := consul.NewRegistry("localhost:8500")

  if err != nil {
    panic(err)
  }
  ctx := context.Background()

  instanceID := discovery.GenerateInstanceID(serviceName)
  if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
    panic(err)
  }

  go func () {
    for {
      if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
        log.Printf("Failed to report healthy state: %v", err)
              }
      time.Sleep(1 * time.Second)
    }
  }()

  defer registry.Deregister(ctx, instanceID, serviceName)

  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  
  grpcServer := grpc.NewServer()
  userUseCase := initUserServer(db)
  handler.NewServer(grpcServer, userUseCase)

  log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
  userRepo := repo.New(db)
  return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {

  err := db.AutoMigrate(&model.User{})
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Migrations run successfully")
  }
}
