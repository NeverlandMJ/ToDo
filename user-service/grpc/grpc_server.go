package grpc

import (
	"context"
	"fmt"
	"log"

	customErr "github.com/NeverlandMJ/ToDo/user-service/pkg/error"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
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
	fmt.Println(req.GetPhone())
	err := g.svc.Otp.CheckOtp(req.GetPhone(), req.GetCode())
	if err != nil {
		if err == customErr.ERR_INCORRECT_CODE {
			return nil, customErr.ERR_INCORRECT_CODE
		}
		log.Println(err)
		return nil, err
	}

	resp, err := g.svc.CreateUserNameAndPassword(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = g.svc.CreateUser(ctx, entity.User{
		Phone: req.Phone,
		UserName: resp.UserName,
		Password: resp.Password,
	})

	if err != nil {
		if err == customErr.ERR_USER_EXIST{
			return nil, customErr.ERR_USER_EXIST
		}
		log.Println(err)
		return nil, err
	}

	return &userpb.ResponseUser{
		Password: resp.Password,
		UserName: resp.UserName,
	}, nil
}




