package controller

import (
	"blog-go/apiv1"
	"blog-go/internal/service"
	"context"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

// Login Login
func (cAuth) Login(ctx context.Context, req *apiv1.LoginReq) (*apiv1.LoginRes, error) {
	login, err := service.Auth.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &apiv1.LoginRes{
		LoginRes: *login,
	}, nil
}

// Register Register
func (cAuth) Register(ctx context.Context, req *apiv1.RegisterReq) (*apiv1.RegisterRes, error) {
	user, err := service.Auth.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return &apiv1.RegisterRes{
		User: *user,
	}, nil
}
