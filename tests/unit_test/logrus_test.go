package unit

import (
	"blog/bootstrap"
	"blog/config"
	"github.com/sirupsen/logrus"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
	bootstrap.Register()
}

func TestLogrus(t *testing.T) {
	logrus.Info("hello")
}
