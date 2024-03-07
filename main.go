package main

import (
	_ "coderblog-interface/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"coderblog-interface/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
