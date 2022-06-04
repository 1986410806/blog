package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpTokenReq struct {
	g.Meta `path:"/qiniu/token" tags:"auth" method:"get" summary:"You first hello api" dc:""`
}

type UpTokenRes struct {
	Token string `json:"token"`
}
