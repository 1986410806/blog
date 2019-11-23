package unit_test

import (
	"blog/bootstrap"
	"blog/config"
	"github.com/sirupsen/logrus"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
}
func TestLogrus(t *testing.T) {
	bootstrap.Register()
	logrus.Info("hello")
}
