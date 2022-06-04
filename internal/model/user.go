package model

import (
	"blog-go/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginRes struct {
	entity.User
	Token      string     `json:"token"`
	ExpireTime gtime.Time `json:"expire_time"`
}

type UserList struct {
	Page
	List []*entity.User
}
