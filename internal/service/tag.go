package service

import (
	"blog-go/apiv1"
	"blog-go/internal/model"
	"blog-go/internal/model/entity"
	"blog-go/internal/service/internal/dao"
	"blog-go/internal/service/internal/do"
	"context"
	"errors"
)

var Tag = sTag{}

type sTag struct{}

var TagNotFound = errors.New("tag not found")

func (s sTag) List(ctx context.Context, req *apiv1.TagListReq) (model.TagListRes, error) {
	var (
		err error
		res = model.TagListRes{
			Page: model.Page{
				Page: req.Page,
				Size: req.Size,
			},
		}
	)

	if res.Total, err = dao.Tag.Total(ctx); err != nil {
		return res, err
	}

	if res.List, err = dao.Tag.List(ctx, req.Page, req.Size); err != nil {
		return res, nil
	}
	return res, err
}

func (s sTag) Del(ctx context.Context, req *apiv1.TagDelReq) error {
	return dao.Tag.Del(ctx, req.Id)
}

func (s sTag) Add(ctx context.Context, req *apiv1.TagAddReq) (*entity.Tag, error) {

	tag, err := dao.Tag.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		tag = &entity.Tag{
			Name: req.Name,
			Desc: req.Desc,
		}
		if err := dao.Tag.Create(ctx, tag); err != nil {
			return nil, err
		}
	}

	return tag, nil
}

func (s sTag) Save(ctx context.Context, req *apiv1.TagSaveReq) error {

	tag, err := dao.Tag.GetById(ctx, req.Id)
	if err != nil {
		return err
	}

	if tag == nil {
		return TagNotFound
	}

	return dao.Tag.Save(ctx, req.Id, do.Tag{
		Name: req.Name,
		Desc: req.Desc,
	})
}
