package user

import (
	"context"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
)

func (s *UserServer) LoginUser(ctx context.Context,
	req *pb.Empty) (*pb.LoginResponse, error) {
	
		//write login logic here
	return &pb.LoginResponse{},nil
}