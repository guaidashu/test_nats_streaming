package main

import (
	"fmt"
	"sync"
	"test_nats_streaming/app"
	"time"
)

func main() {
	// var (
	// 	flagArr []string
	// )
	//
	// flag.Parse()
	//
	// var num = 300
	// flagArr = flag.Args()
	// if len(flagArr) >= 1 {
	// 	if len(flagArr) >= 2 {
	// 		num, _ = strconv.Atoi(flagArr[1])
	// 	}
	// 	// 进入选择
	// 	switch flagArr[0] {
	// 	case "pub":
	// 		// 启动推送
	// 		waitGroup := sync.WaitGroup{}
	// 		startTime := time.Now().UnixNano()
	// 		for i := 0; i < 100; i++ {
	// 			waitGroup.Add(1)
	// 			go func() {
	// 				app.Publisher(num)
	// 				waitGroup.Done()
	// 			}()
	// 		}
	// 		waitGroup.Wait()
	// 		endTime := time.Now().UnixNano()
	// 		result := float64(endTime-startTime) / 1000000000.0
	// 		fmt.Println(result)
	// 	case "pub_async":
	// 		app.PublishAsync(num)
	// 	case "sub":
	// 		app.Subscribe()
	// 	}
	// } else {
	// 	fmt.Println("请输入正确的参数")
	// }

	go func() {
		app.Subscribe()
	}()

	for {
		var num int
		_, _ = fmt.Scanln(&num)
		// 启动推送
		waitGroup := sync.WaitGroup{}
		startTime := time.Now().UnixNano()
		for i := 0; i < 100; i++ {
			waitGroup.Add(1)
			go func() {
				app.Publisher(num)
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
		endTime := time.Now().UnixNano()
		result := float64(endTime-startTime) / 1000000000.0
		fmt.Println(result)
	}
}
