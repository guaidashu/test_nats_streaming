package main

import (
	"flag"
	"fmt"
	"test_nats_streaming/app"
)

func main() {
	var (
		flagArr []string
	)

	flag.Parse()

	flagArr = flag.Args()
	if len(flagArr) >= 1 {
		// 进入选择
		switch flagArr[0] {
		case "pub":
			app.Publisher()
		case "sub":
			app.Subscribe()
		}
	} else {
		fmt.Println("请输入正确的参数")
	}
}
