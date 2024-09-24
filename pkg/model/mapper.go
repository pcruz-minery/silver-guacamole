package model

import (
  "github.com/mineryreport/flexicar/gen"
)

func UserToProto(u *User) *gen.User {
  return &gen.User{
    Id: u.ID,
    Name: u.Name,
    Email: u.Email,
  }
}

func UserFromProto(u *gen.UserProfileResponse) *User {
  return &User{
    ID: u.Id,
    Name: u.Name,
    Email: u.Email,
  }
}

