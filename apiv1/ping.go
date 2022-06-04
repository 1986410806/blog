package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PingReq struct {
	g.Meta `path:"/ping" tags:"ping" method:"get" summary:"You first hello api"`
}
type PingRes struct{}
