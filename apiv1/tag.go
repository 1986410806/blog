package apiv1

import (
	"blog-go/internal/model"
	"blog-go/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type TagListReq struct {
	g.Meta `path:"/tags" tags:"tag" method:"get" summary:"You first hello api"`
	CommonPaginationReq
}
type TagListRes struct {
	model.TagListRes
}

type TagAddReq struct {
	g.Meta `path:"/tag/add" tags:"tag" method:"post" summary:"You first hello api"`
	Name   string `q:"name" v:"required"`
	Desc   string `q:"desc" v:"required"`
}
type TagAddRes struct {
	entity.Tag
}

type TagSaveReq struct {
	g.Meta `path:"/tag/save" tags:"tag" method:"post" summary:"You first hello api"`
	Id     int    `q:"id" v:"required"`
	Name   string `q:"name" v:"required"`
	Desc   string `q:"desc" v:"required"`
}

type TagSaveRes struct {
}

type TagDelReq struct {
	g.Meta `path:"/tag/del" tags:"tag" method:"post" summary:"You first hello api"`
	Id     int `q:"id" v:"required"`
}
type TagDelRes struct {
}
