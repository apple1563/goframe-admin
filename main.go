package main

import (
	_ "goframe-starter/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-starter/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
