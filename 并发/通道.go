package main

import "fmt"

//通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

func main() {
	//使用make创建通道
	messages := make(chan string) //无缓冲通道

	//使用 channel <- 语法发送值到通道中
	go func() {
		messages <- "hello"
	}()

	//使用<-从通道中接受值
	msg := <-messages
	fmt.Println(msg)

	//默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。这个特性允许我们，不使用任何其它的同步操作，就可以在程序结尾处等待消息。

	mess := make(chan string, 4)

	mess <- "hello world"
	mess <- "hello wuhan"

	msg = <-mess
	msg = <-mess
	fmt.Println(msg)
	//fmt.Println(msg)
	fmt.Println(<-mess)
}
