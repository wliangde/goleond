/**
User:       wliangde
CreateTime: 18/3/1 上午10:37
**/
package test

import (
	"fmt"
	"os"
	"time"
)

func CountDown() {
	abort := make(chan struct{})
	go func() {
		fmt.Println("监听输入")
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
		close(abort)
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := 10; i > 0; i-- {
		select {
		//case <-time.Tick(time.Second): //会有goroutine泄漏
		case <-ticker.C:
			fmt.Println(i)
		case <-abort:
			fmt.Println("结束倒计时")
			return
		}
	}
}
