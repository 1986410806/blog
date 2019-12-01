package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *config

type config struct {
	Env        string `yaml:"Env"`        // 环境：prod、dev、local
	BaseUrl    string `yaml:"BaseUrl"`    // base url
	Port       string `yaml:"Port"`       // 端口
	LogFile    string `yaml:"LogFile"`    // 日志文件
	Debug      bool   `yaml:"Debug"`      // 是否显示日志
	StaticPath string `yaml:"StaticPath"` // 静态文件目录

	MySqlUrl string `yaml:"MySqlUrl"` // 数据库连接地址

	Redis struct {
		Host string `yaml:"Host"`
		Port int    `yaml:"Port"`
		DB   int    `yaml:"DB"`
	} `yaml:"Redis"`
	// jwt 加密密匙
	JwtSecret string `yaml:"JwtSecret"`
	// Github
	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	// QQ登录
	QQConnect struct {
		AppId  string `yaml:"AppId"`
		AppKey string `yaml:"AppKey"`
	} `yaml:"QQConnect"`

	// 七牛云oss配置
	QiniuOss struct {
		Host      string `yaml:"Host"`
		Bucket    string `yaml:"Bucket"`
		Endpoint  string `yaml:"Endpoint"`
		AccessKey string `yaml:"AccessKey"`
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"QiniuOss"`
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
		return
	}

	Conf = &config{}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		panic(err)
	}
}
