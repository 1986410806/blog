// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-go/internal/model/entity"
	"blog-go/internal/service/internal/dao/internal"
	"blog-go/internal/service/internal/do"
	"context"
)

// tagDao is the data access object for table tag.
// You can define custom methods on it to extend its functionality as you wish.
type tagDao struct {
	*internal.TagDao
}

var (
	// Tag is globally public accessible object for table tag operations.
	Tag = tagDao{
		internal.NewTagDao(),
	}
)

// Fill with you ideas below.

func (d tagDao) GetByName(ctx context.Context, name string) (*entity.Tag, error) {
	var tag *entity.Tag
	if err := d.Ctx(ctx).Where(do.Tag{Name: name}).Scan(&tag); err != nil {
		return nil, err
	}
	return tag, nil
}
func (d tagDao) GetById(ctx context.Context, id int) (*entity.Tag, error) {
	var tag *entity.Tag
	if err := d.Ctx(ctx).Where(do.Tag{Id: id}).Scan(&tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (d tagDao) Create(ctx context.Context, user *entity.Tag) error {
	_, err := d.Ctx(ctx).Insert(user)
	return err
}

func (d tagDao) Del(ctx context.Context, id int) error {
	_, err := d.Ctx(ctx).Delete("id", id)
	return err
}

func (d tagDao) List(ctx context.Context, page, size int) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	err := d.Ctx(ctx).Page(page, size).Scan(&tags)
	return tags, err
}

//Total page Total
func (d tagDao) Total(ctx context.Context) (int, error) {
	return d.Ctx(ctx).Count()
}

func (d tagDao) Save(ctx context.Context, id int, tag do.Tag) error {
	_, err := d.Ctx(ctx).Where("id", id).Save(tag)
	return err
}
