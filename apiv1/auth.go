package apiv1

import (
	"blog-go/internal/model"
	"blog-go/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginReq LoginReq
type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"auth" method:"post" summary:"You first hello api"`
	Email    string `p:"email" v:"required|email"`
	Password string `p:"password" v:"required|password"`
	Type     uint8  `p:"type" v:"required" desc:"1 后台 2 blog user"`
}
type LoginRes struct {
	model.LoginRes
}
type RegisterReq struct {
	g.Meta    `path:"/auth/register" tags:"auth" method:"post" summary:"You first hello api" dc:""`
	Avatar    string `p:"avatar"  v:"required"   `       //
	Email     string `p:"email"    v:"required|email"  ` //
	Name      string `p:"name"    v:"required"   `       //
	Password  string `p:"password" v:"required|password"`
	Type      int    `p:"type"     v:"required"  ` //
	Introduce string `p:"introduce"  v:"required"` //
}
type RegisterRes struct {
	entity.User
}

type CurrentReq struct {
	g.Meta `path:"/user/current" tags:"auth" method:"get" summary:"You first hello api" dc:""`
}

type CurrentRes struct {
	entity.User
}

type UserDelReq struct {
	g.Meta `path:"/user/del" tags:"auth" method:"post" summary:"You first hello api" dc:""`
	Uid    int `p:"uid" v:"required"`
}
type UserDelRes struct {
}
type UserListReq struct {
	g.Meta `path:"/user/list" tags:"auth" method:"get" summary:"You first hello api" dc:""`
	CommonPaginationReq
}
type UserListRes struct {
	model.UserList
}
