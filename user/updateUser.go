package user

import (
	"context"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
)

func (s *UserServer) UpdateUser(ctx context.Context,
	 req *pb.UpdateUserRequest) (*pb.User, error){

		// write logic here
		return &pb.User{},nil
	 }