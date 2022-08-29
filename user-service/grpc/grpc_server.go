package grpc

import (
	"context"
	"errors"
	"log"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	customErr "github.com/NeverlandMJ/ToDo/user-service/pkg/error"
	"github.com/NeverlandMJ/ToDo/user-service/service"
	"github.com/NeverlandMJ/ToDo/user-service/v1/userpb"
)

type gRPCServer struct {
	userpb.UnimplementedUserServiceServer
	svc service.Service
}

func NewgRPCServer(svc service.Service) *gRPCServer {
	return &gRPCServer{
		svc: svc,
	}
}

func (g *gRPCServer) SendCode(ctx context.Context, req *userpb.RequestPhone) (*userpb.RequestPhone, error) {
	phone := req.GetPhone()
	_, err := g.svc.Otp.SendOtp(phone)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &userpb.RequestPhone{
		Phone: phone,
	}, nil

}

func (g *gRPCServer) RegisterUser(ctx context.Context, req *userpb.Code) (*userpb.ResponseUser, error) {
	err := g.svc.Otp.CheckOtp(req.GetPhone(), req.GetCode())
	if err != nil {
		if err == customErr.ERR_INCORRECT_CODE {
			return nil, customErr.ERR_INCORRECT_CODE
		} else {
			log.Println(err)
			return nil, err
		}
	}

	resp, err := g.svc.CreateUsernameAndPassword(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = g.svc.CreateUser(ctx, entity.User{
		Phone:    req.Phone,
		UserName: resp.UserName,
		Password: resp.Password,
	})

	if err != nil {
		if errors.Is(err, customErr.ERR_USER_EXIST) {
			log.Printf("user exist %v\n", err)
			return nil, customErr.ERR_USER_EXIST
		} else {
			log.Printf("other error %v\n", err)
			return nil, err
		}
	}

	return &userpb.ResponseUser{
		Password: resp.Password,
		UserName: resp.UserName,
	}, nil
}

func (g *gRPCServer) SignIn(ctx context.Context, req *userpb.SignInUer) (*userpb.User, error) {
	un := req.GetUserName()
	pw := req.GetPassword()

	user, err := g.svc.GetUser(ctx, un, pw)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &userpb.User{
		ID:        user.ID,
		UserName:  user.UserName,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		IsBlocked: user.IsBlocked,
	}, nil
}
