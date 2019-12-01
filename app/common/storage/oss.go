package storage

import (
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"qiniupkg.com/x/bytes.v7"
	"time"

	"blog/config"
	"github.com/go-resty/resty/v2"
	"github.com/mlogclub/simple"
)

// 上传图片
func PutImage(data []byte) (string, error) {
	md5 := simple.MD5Bytes(data)
	key := "images/" + simple.TimeFormat(time.Now(), "2006/01/02/") + md5 + ".jpg"
	return putObject(key, data)
}

// 上传
func putObject(key string, data []byte) (string, error) {
	upToken := GetUpToken()
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei, // 空间对应的机房
		UseHTTPS:      false,               // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), int64(len(data)), &putExtra)
	return ret.Key, err

}

// 将图片copy到oss
func CopyImage(inputUrl string) (string, error) {
	data, err := download(inputUrl)
	if err != nil {
		return "", err
	}
	return PutImage(data)
}

// 下载
func download(url string) ([]byte, error) {
	rsp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}
	return rsp.Body(), nil
}

// 获取上传令牌
func GetUpToken() string {
	putPolicy := storage.PutPolicy{
		Scope:      config.Conf.QiniuOss.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    7200, //示例2小时有效期
	}

	mac := qbox.NewMac(config.Conf.QiniuOss.AccessKey, config.Conf.QiniuOss.SecretKey)

	return putPolicy.UploadToken(mac)
}
