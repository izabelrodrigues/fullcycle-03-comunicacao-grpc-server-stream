package services

import (
	"context"
	"fmt"
	"time"

	"github.com/izabelrodrigues/fullcycle-grpc-stream/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	//Insert database
	fmt.Println(req.Nome)

	return &pb.User{
		Id: "123",
		Nome: req.GetNome(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) (error) {
	stream.Send(&pb.UserResultStream {
		Status: "Init",
		User: &pb.User{},
	})

	time.Sleep(time.Second *3)

	stream.Send(&pb.UserResultStream {
		Status: "Inserting",
		User: &pb.User{},
	})

	time.Sleep(time.Second *3)

	stream.Send(&pb.UserResultStream {
		Status: "User has been inserted",
		User: &pb.User{
			Id: "124",
			Nome: req.GetNome(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second *3)

	stream.Send(&pb.UserResultStream {
		Status: "Completed",
		User: &pb.User{
			Id: "124",
			Nome: req.GetNome(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second *3)

	return nil
}