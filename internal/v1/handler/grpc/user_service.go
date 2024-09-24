package grpc

import (
	"context"
	"errors"
	"log"

	pb "github.com/mineryreport/flexicar/gen"
	interfaces "github.com/mineryreport/flexicar/internal/v1"
	"github.com/mineryreport/flexicar/pkg/model"
	"google.golang.org/grpc"
)

type UserServStruct struct {
  
  useCase interfaces.UseCaseInterface
  pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
  userGrpc := &UserServStruct{useCase: usecase}
  pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (srv *UserServStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {

  data := srv.transformUserRPC(req)
  if data.Email == "" || data.Name == "" {
    return &pb.UserProfileResponse{}, errors.New("Name and Email cannot be empty")
  }
  user, err := srv.useCase.Create(data)
  if err != nil {
    return &pb.UserProfileResponse{}, err
  }
  return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
  id := req.GetId()
  if id == "" {
    return &pb.UserProfileResponse{}, errors.New("Id cannot be empty")
      }
  user, err := srv.useCase.Get(id)
  if err != nil {
    return &pb.UserProfileResponse{}, err
      }
  return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.SuccessResponse, error) {
  log.Fatalf("Non Implemented Method User.Update")
  return &pb.SuccessResponse{
  }, nil
}




func (srv *UserServStruct) transformUserRPC(req *pb.CreateUserRequest) model.User {
  return model.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServStruct) transformUserModel(user model.User) *pb.UserProfileResponse {
  return &pb.UserProfileResponse{Id: string(user.ID), Name: user.Name, Email: user.Email}
}


