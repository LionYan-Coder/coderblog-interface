package main

import (
	_ "coderblog-interface/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "coderblog-interface/internal/logic"

	"coderblog-interface/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
