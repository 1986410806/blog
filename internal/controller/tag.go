package controller

import (
	"blog-go/apiv1"
	"blog-go/internal/service"
	"context"
)

var Tag = cTag{}

type cTag struct{}

func (cTag) List(ctx context.Context, req *apiv1.TagListReq) (*apiv1.TagListRes, error) {
	list, err := service.Tag.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return &apiv1.TagListRes{
		TagListRes: list,
	}, nil
}
func (cTag) Add(ctx context.Context, req *apiv1.TagAddReq) (*apiv1.TagAddRes, error) {
	tag, err := service.Tag.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return &apiv1.TagAddRes{
		Tag: *tag,
	}, nil
}
func (cTag) Save(ctx context.Context, req *apiv1.TagSaveReq) (res *apiv1.TagSaveRes, err error) {
	err = service.Tag.Save(ctx, req)

	return
}
func (cTag) Del(ctx context.Context, req *apiv1.TagDelReq) (res *apiv1.TagDelRes, err error) {

	err = service.Tag.Del(ctx, req)
	return
}
