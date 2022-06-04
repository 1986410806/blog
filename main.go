package main

import (
	_ "blog-go/internal/packed"

	"blog-go/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
