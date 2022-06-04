package controller

import (
	"blog-go/apiv1"
	"blog-go/internal/service"
	"context"
)

var (
	File = cFile{}
)

type cFile struct{}

func (*cFile) QiniuToken(ctx context.Context, _ *apiv1.UpTokenReq) (*apiv1.UpTokenRes, error) {
	var token = service.GetUpToken(ctx)
	return &apiv1.UpTokenRes{
		Token: token,
	}, nil
}
