package main

import (
	"flag"
	"fmt"

	"github.com/liuxiaobopro/go-lib/console"
)

type options struct {
	GenerateApi string
	Help        bool
}

var op = new(options)

func (th *options) parseOptions() {
	flag.StringVar(&th.GenerateApi, "gen_api", "", "请输入要生成的api(admin/api)")
	flag.BoolVar(&th.Help, "help", true, "")

	flag.Parse()
}

func main() {
	// 获取flag参数
	op.parseOptions()

	if op.GenerateApi != "" {
		console.Console.Info("开始生成api:" + op.GenerateApi)
	}

	if op.Help {
		fmt.Print(`
		
		import (
			"flag"
			"log"
		
			"github.com/liuxiaobopro/go-lib/console"
		)

		`)
	}
}
