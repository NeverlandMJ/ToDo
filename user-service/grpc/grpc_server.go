package grpc

import (
	"context"
	"errors"
	"log"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/customErr"
	"github.com/NeverlandMJ/ToDo/user-service/service"
	"github.com/NeverlandMJ/ToDo/user-service/userpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.Internal, "error occured while sending TOTP")
	}

	return &userpb.RequestPhone{
		Phone: phone,
	}, nil

}

func (g *gRPCServer) RegisterUser(ctx context.Context, req *userpb.Code) (*userpb.ResponseUser, error) {
	err := g.svc.Otp.CheckOtp(req.GetPhone(), req.GetCode())
	if err != nil {
		if errors.Is(err, customErr.ERR_INCORRECT_CODE) {
			return nil, status.Error(codes.Unauthenticated, customErr.ERR_INCORRECT_CODE.Error())
		} else {
			log.Println(err)
			return nil, status.Error(codes.Internal, "error occured while checking TOTP")
		}
	}

	resp, err := g.svc.CreateUsernameAndPassword(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "error occured while creating username and password")
	}

	err = g.svc.CreateUser(ctx, entity.User{
		Phone:    req.Phone,
		UserName: resp.UserName,
		Password: resp.Password,
	})

	if err != nil {
		if errors.Is(err, customErr.ERR_USER_EXIST) {
			log.Printf("user exist %v\n", err)
			return nil, status.Error(codes.AlreadyExists, customErr.ERR_USER_EXIST.Error())
		} else {
			log.Printf("other error %v\n", err)
			return nil, status.Error(codes.Internal, "error occured while creating user")
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
		if errors.Is(err, customErr.ERR_USER_NOT_EXIST) {
			return nil, status.Error(codes.NotFound, customErr.ERR_USER_NOT_EXIST.Error())
		} else if errors.Is(err, customErr.ERR_INCORRECT_PASSWORD) {
			return nil, status.Error(codes.Unauthenticated, customErr.ERR_INCORRECT_PASSWORD.Error())
		} else {
			return nil, status.Error(codes.Internal, "unexpected error")
		}
	}

	return &userpb.User{
		ID:        user.ID.String(),
		UserName:  user.UserName,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		IsBlocked: user.IsBlocked,
	}, nil
}

func (g *gRPCServer) ChangePassword(ctx context.Context, req *userpb.RequestChangePassword) (*userpb.Empty, error) {
	id, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.ChangePassword(ctx, id, req.OldPassword, req.NewPassword)
	if err != nil {
		log.Println(err)
		if errors.Is(err, customErr.ERR_USER_EXIST) {
			return nil, status.Error(codes.NotFound, "user with given id not found")
		} else if errors.Is(err, customErr.ERR_INCORRECT_PASSWORD) {
			return nil, status.Error(codes.PermissionDenied, "old password is incorrect")
		}else {
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &userpb.Empty{}, nil
}

func (g *gRPCServer) ChangeUserName(ctx context.Context, req *userpb.RequestUserName) (*userpb.Empty, error) {
	id, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.ChangeUserName(ctx, id, req.UserName)
	if err != nil {
			log.Println(err)
		if errors.Is(err, customErr.ERR_USER_EXIST) {
			return nil, status.Error(codes.NotFound, "user with given id not found")
		}else {
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &userpb.Empty{}, nil
}

func (g *gRPCServer) DeleteAccount(ctx context.Context, req *userpb.RequestDeleteAccount) (*userpb.Empty, error)  {
	id, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.DeleteAccount(ctx, id, req.Password, req.UserName)
	if err != nil {
			log.Println(err)
		if errors.Is(err, customErr.ERR_USER_EXIST) {
			return nil, status.Error(codes.NotFound, "user with given id not found")
		}else if errors.Is(err, customErr.ERR_INCORRECT_PASSWORD){
			return nil, status.Error(codes.PermissionDenied, "password is incorrect")
		}else {
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &userpb.Empty{}, nil
}