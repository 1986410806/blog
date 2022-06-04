package controller

import (
	"blog-go/apiv1"
	"blog-go/internal/service"
	"context"
)

var (
	User = cUser{}
)

type cUser struct{}

func (cUser) List(ctx context.Context, req *apiv1.UserListReq) (*apiv1.UserListRes, error) {
	users, err := service.Auth.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return &apiv1.UserListRes{
		UserList: users,
	}, nil
}

func (cUser) Del(ctx context.Context, req *apiv1.UserDelReq) (res *apiv1.UserDelRes, err error) {
	err = service.Auth.Del(ctx, req)

	return
}

func (cUser) Current(ctx context.Context, _ *apiv1.CurrentReq) (*apiv1.CurrentRes, error) {
	user, err := service.Auth.Current(ctx)
	if err != nil {
		return nil, err
	}
	return &apiv1.CurrentRes{
		User: *user,
	}, err
}
