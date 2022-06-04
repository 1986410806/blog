package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	bucket    = "zhao-blog"
	accessKey = g.Cfg().MustGet(context.TODO(), "qiniu.key").String()
	secretKey = g.Cfg().MustGet(context.TODO(), "qiniu.secret").String()
)

func GetUpToken(_ context.Context) string {
	var token string

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	token = putPolicy.UploadToken(qbox.NewMac(accessKey, secretKey))

	return token
}
