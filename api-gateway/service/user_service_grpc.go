package service

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/auth"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
	customErr "github.com/NeverlandMJ/ToDo/api-gateway/pkg/error"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/utilities"
	"github.com/NeverlandMJ/ToDo/api-gateway/v1/userpb"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
		client:     client,
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

func (c userServiceGRPCClient) RegisterUser(ctx context.Context, code entity.ReqSignUp) (entity.RespUser, error) {
	phone, err := c.inMemoryDB.Get(code.Phone).Result()
	if err != nil && phone == "" {
		return entity.RespUser{}, customErr.ERR_CODE_HAS_EXPIRED
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

func (c userServiceGRPCClient) SignIn(ctx context.Context, data entity.ReqSignIn) (string, error) {
	resp, err := c.client.SignIn(ctx, &userpb.SignInUer{
		UserName: data.UserName,
		Password: data.Password,
	})
	if err != nil {
		return "", err
	}

	if resp.IsBlocked {
		return "", customErr.ERR_USER_BLOCKED
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &auth.Claims{
		ID:          resp.GetID(),
		PhoneNumber: resp.GetPhone(),
		IsBlocked:   resp.GetIsBlocked(),
		UserName:    resp.GetUserName(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(auth.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", customErr.ERR_INTERNAL
	}

	return tokenString, nil
}

func (c userServiceGRPCClient) ChangePassword(ctx context.Context, userID string, new entity.ReqChangePassword) error {
	_, err := c.client.ChangePassword(ctx, &userpb.RequestChangePassword{
		UserID:      userID,
		OldPassword: new.OldPassword,
		NewPassword: new.NewPassword,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c userServiceGRPCClient) ChangeUserName(ctx context.Context, userID string, new entity.ReqChangeUsername) error {
	_, err := c.client.ChangeUserName(ctx, &userpb.RequestUserName{
		UserID:   userID,
		UserName: new.UserName,
	})

	if err != nil {
		return err
	}
	return nil
}

func (c userServiceGRPCClient) DeleteAccount(ctx context.Context, userID string, auth entity.ReqSignIn) error {
	_, err := c.client.DeleteAccount(ctx, &userpb.RequestDeleteAccount{
		UserID:   userID,
		Password: auth.Password,
		UserName: auth.UserName,
	})

	if err != nil {
		return err
	}
	return nil
}
