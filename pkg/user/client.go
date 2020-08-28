package user

import (
	"context"
	"github.com/jsanda/user-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	connection *grpc.ClientConn
	client pb.UserServiceClient
}

type User struct {
	Email string
	Name string
}

func NewUserServiceClient(addr string) (*ServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, err
	}

	return &ServiceClient{connection: conn, client: pb.NewUserServiceClient(conn)}, nil
}

func (c *ServiceClient) Close() error {
	return c.connection.Close()
}

func (c *ServiceClient) CreateUser(ctx context.Context, user User) error {
	pbUser := &pb.User{Email: user.Email, Name: user.Name}
	_, err := c.client.CreateUser(ctx, pbUser)

	return err
}

func (c *ServiceClient) GetUsers(ctx context.Context) ([]User, error) {
	response, err := c.client.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]User, len(response.Users))

	for i, u := range response.Users {
		users[i] = User{Email: u.Email, Name: u.Name}
	}

	return users, nil
}
