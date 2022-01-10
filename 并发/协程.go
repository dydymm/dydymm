package main

import (
	"fmt"
	"time"
)

//协程,是轻量级的执行线程。

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("xiecheng1") //go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会并发地执行这个函数。

	//为匿名函数启动一个协程
	go func(msg string) {
		fmt.Println(msg)
	}("niming")
	//现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成
	time.Sleep(time.Second)
	fmt.Println(("finish"))

}
