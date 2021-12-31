package main

import (
	"fmt"
)

func main() {
	//var a int = 40
	//b := 60

	/*if a < 30 {//当a<30时，都打印
		fmt.Println("a<30")
	}
	fmt.Println("a>30")*/

	/*if a > 30 {//当a>30时，只打印a>30
		fmt.Println("a>30")
	} else {
		fmt.Println("a<30")
	}*/

	/*if a == 40 {
		if b == 80 {
			fmt.Printf("a=%d,b=%d", a, b)
		} else if b == 60 {
			fmt.Printf("a=%d,b不是80，为%d", a, b)
		}
	}*/

	/*
		switch a {
		case 40:
			fmt.Printf("b=%d", b)

		}*/
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello" //c1,c2随机执行。
	c1 <- "world"
	select { //仅能用于 信道/通道 的相关操作
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2: //在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有一个信道有接收到数据，那么 select 就结束，所以输出如下
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}
}
