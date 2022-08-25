package service

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/api-gateway/entity"
	"github.com/NeverlandMJ/ToDo/api-gateway/v1/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userServiceGRPCClient struct {
	client userpb.UserServiceClient
}

func NewGRPCClientUser(url string) userServiceGRPCClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client := userpb.NewUserServiceClient(conn)
	return userServiceGRPCClient{
		client: client,
	}
}

func (c userServiceGRPCClient) SendCode(ctx context.Context, ph entity.ReqPhone) (entity.ReqPhone, error){
	resp, err := c.client.SendCode(ctx, &userpb.RequestPhone{
		Phone: ph.Phone,
	})
	if err != nil {
		return entity.ReqPhone{}, err
	}

	return entity.ReqPhone{
		Phone: resp.Phone,
	}, nil
}

func (c userServiceGRPCClient) RegisterUser(ctx context.Context, code entity.ReqCode) (entity.RespUser, error)  {
	resp, err := c.client.RegisterUser(ctx, &userpb.Code{
		Phone: code.Phone,
		Code: code.Code,
	})
	if err != nil {
		return entity.RespUser{}, err
	}

	return entity.RespUser{
		UserName: resp.GetUserName(),
		Password: resp.GetPassword(),
	}, nil
}