package controller

import (
	"context"

	"blog-go/apiv1"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, _ *apiv1.PingReq) (res *apiv1.PingRes, err error) {
	g.RequestFromCtx(ctx).Response.WritelnExit("pong")
	return
}
