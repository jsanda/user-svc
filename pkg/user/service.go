package user

import (
	"context"
	"github.com/jsanda/user-svc/pkg/cassandra"
	"github.com/jsanda/user-svc/pkg/pb"
)

type UserService struct {
	cassClient *cassandra.Client
}

func NewUserService(cassandraSvc string) (*UserService, error) {
	cassClient, err := cassandra.NewClient(cassandraSvc)
	if err != nil {
		return nil, err
	}

	if err = cassClient.InitSchema(); err != nil {
		return nil, err
	}

	return &UserService{cassClient: cassClient}, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	err := s.cassClient.CreateUser(cassandra.User{Email: user.Email, Name: user.Name})
	return &pb.CreateUserResponse{}, err
}

func (s *UserService) GetUsers(context.Context, *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := s.cassClient.GetUsers()
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.User, 0)
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{Email: u.Email, Name: u.Name})
	}

	return &pb.GetUsersResponse{Users: pbUsers}, nil
}
