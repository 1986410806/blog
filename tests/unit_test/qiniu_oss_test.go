package unit

import (
	"blog/app/common/storage"
	"blog/bootstrap"
	"blog/config"
	"bufio"
	"fmt"
	"os"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
	bootstrap.Register()
}

// 获取上传token
func TestUpToken(t *testing.T) {
	token := storage.GetUpToken()
	fmt.Println(token)
}

func TestPutFile(t *testing.T) {
	localFile := "/Users/zhaohuinan/Code/study/golang/blog/storage/static/aaa.png"
	file, err := os.Open(localFile)

	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		t.Error(statsErr)
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)

	bufr.Read(bytes)

	str, err := storage.PutImage(bytes)
	if err != nil {
		t.Error(err)
	}
	println(str)
}
