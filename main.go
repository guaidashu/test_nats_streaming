package main

import (
	"flag"
	"fmt"
	"strconv"
	"test_nats_streaming/app"
)

func main() {
	var (
		flagArr []string
	)

	flag.Parse()

	var num = 300
	flagArr = flag.Args()
	if len(flagArr) >= 1 {
		if len(flagArr) >= 2 {
			num, _ = strconv.Atoi(flagArr[1])
		}
		// 进入选择
		switch flagArr[0] {
		case "pub":
			app.Publisher(num)
		case "pub_async":
			app.PublishAsync(num)
		case "sub":
			app.Subscribe()
		}
	} else {
		fmt.Println("请输入正确的参数")
	}
}
