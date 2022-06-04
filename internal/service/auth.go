package service

import (
	"blog-go/apiv1"
	"blog-go/internal/model"
	"blog-go/internal/model/entity"
	"blog-go/internal/service/internal/dao"
	"blog-go/utility"
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var Auth = auth{}

type auth struct{}

var UserNotFound = errors.New("user not found")

//Login Login
func (s auth) Login(ctx context.Context, req *apiv1.LoginReq) (*model.LoginRes, error) {
	user, err := dao.User.GetUserByEmail(ctx, req.Email)

	switch {
	case err != nil:
		return nil, err
	case user == nil:
		return nil, errors.New("账户不存在")
	}

	if md5 := gmd5.MustEncryptBytes([]byte(req.Password)); md5 != user.Password {
		return nil, errors.New("密码错误")
	}
	secret, err := g.Cfg().Get(ctx, "token.secret")
	if err != nil {
		return nil, err
	}

	token, err := utility.NewJwtToken(secret.Bytes(), user.Id)
	if err != nil {
		return nil, err
	}

	return &model.LoginRes{
		User:       *user,
		Token:      token.Token,
		ExpireTime: *gtime.New(token.Exp),
	}, err
}

//Register Register
func (s auth) Register(ctx context.Context, req *apiv1.RegisterReq) (*entity.User, error) {
	user, err := dao.User.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("账户已存在")
	}

	user = &entity.User{
		Email:     req.Email,
		Name:      req.Name,
		Password:  gmd5.MustEncryptBytes([]byte(req.Password)),
		Type:      0,
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
		Introduce: "",
	}

	if err := dao.User.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, err
}

func (s auth) List(ctx context.Context, req *apiv1.UserListReq) (model.UserList, error) {
	var (
		err error
		res = model.UserList{
			Page: model.Page{
				Page: req.Page,
				Size: req.Size,
			},
		}
	)

	if res.Total, err = dao.User.Total(ctx); err != nil {
		return res, err
	}

	if res.List, err = dao.User.List(ctx, req.Page, req.Size); err != nil {
		return res, nil
	}
	return res, err
}

func (s auth) Del(ctx context.Context, req *apiv1.UserDelReq) error {

	user, err := s.getUser(ctx, req.Uid)
	if err != nil {
		return err
	}
	return dao.User.Del(ctx, user)
}

func (s auth) getUser(ctx context.Context, id int) (*entity.User, error) {
	user, err := dao.User.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, UserNotFound
	}
	return user, nil
}

func (s auth) Current(ctx context.Context) (*entity.User, error) {
	if uid := g.RequestFromCtx(ctx).Get("uid").Int(); uid != 0 {
		return s.getUser(ctx, uid)
	}
	return nil, UserNotFound
}
