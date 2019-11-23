package main

import (
	"blog/bootstrap"
	"blog/config"
	"flag"
)

var configFile = flag.String("config", "./blog.yaml", "配置文件路径")

func init() {

	flag.Parse()

	config.InitConfig(*configFile)
}

func main() {
	// 运行应用
	app := bootstrap.Register()

	err := bootstrap.Run(app)
	if err != nil {
		panic(err)
	}
}
