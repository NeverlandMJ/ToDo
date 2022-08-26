package service

import (
	"context"
	"fmt"
	"time"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/utilities"
	"github.com/NeverlandMJ/ToDo/api-gateway/v1/userpb"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ERR_CODE_HAS_EXPIRED = fmt.Errorf("code has been expired")

type userServiceGRPCClient struct {
	client     userpb.UserServiceClient
	inMemoryDB *redis.Client
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

	db, err := utilities.NewRedisClient()
	if err != nil {
		panic(err)
	}

	return userServiceGRPCClient{
		client: client,
		inMemoryDB: db,
	}
}

func (c userServiceGRPCClient) SendCode(ctx context.Context, ph entity.ReqPhone) (entity.ReqPhone, error) {
	resp, err := c.client.SendCode(ctx, &userpb.RequestPhone{
		Phone: ph.Phone,
	})
	if err != nil {
		return entity.ReqPhone{}, err
	}
	err = c.inMemoryDB.Set(ph.Phone, ph.Phone, time.Minute).Err()
	if err != nil {
		return entity.ReqPhone{}, err
	}
	return entity.ReqPhone{
		Phone: resp.Phone,
	}, nil
}

func (c userServiceGRPCClient) RegisterUser(ctx context.Context, code entity.ReqCode) (entity.RespUser, error) {
	phone, err := c.inMemoryDB.Get(code.Phone).Result()
	if err != nil && phone == "" {
		return entity.RespUser{}, ERR_CODE_HAS_EXPIRED
	}
	
	code.Phone = phone
	resp, err := c.client.RegisterUser(ctx, &userpb.Code{
		Phone: code.Phone,
		Code:  code.Code,
	})
	if err != nil {
		return entity.RespUser{}, err
	}

	return entity.RespUser{
		UserName: resp.GetUserName(),
		Password: resp.GetPassword(),
	}, nil
}
