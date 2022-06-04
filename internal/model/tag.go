package model

import "blog-go/internal/model/entity"

type TagListRes struct {
	Page
	List []*entity.Tag
}
