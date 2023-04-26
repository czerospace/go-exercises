package main

import (
	"fmt"
	"time"
)

// 定义两个 channel

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		// 读取 chan 等待 另一个 goroutine 的通知打印数字
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		// 通知另一个 goroutine 去打印字母
		letter <- true
	}
}

func printLtter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		// 读取 chan 等待 另一个 goroutine 的通知打印数字
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		// 通知另一个 goroutine 去打印字母
		number <- true
	}
}

/*
	使用两个 goroutine 交替打印序列，一个 goroutine 打印数字，另一个 goroutine 打印字母
	效果如下：
	12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

*/

func main() {
	go printNum()
	go printLtter()
	number <- true

	time.Sleep(100 * time.Second)
}
