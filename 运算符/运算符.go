package main

import (
	"fmt"
)

func main() {
	/*//算数运算符
	a := 24
	b := 10
	var c int

	c = a + b
	fmt.Printf("第一行c=%d\n", c)
	c = a / b
	fmt.Printf("第二行c=%d\n", c)
	c = a % b
	fmt.Printf("第三行c=%d\n", c)
	a = 30
	a++
	fmt.Printf("第四行c=%d\n", a)
	a = 30
	a--
	fmt.Printf("第五行c=%d\n", a)
	*/
	/*//关系运算符 == != < > <=
	a := 10
	b := 10
	//var c int
	if a != b {
		fmt.Println("a=b")
	} else {
		fmt.Println("a!=b")
	}
	*/
	/*//逻辑运算符
	a := true
	b := false
	if a && b {
		fmt.Println("第一行为，true!")
	}
	if a || b {
		fmt.Println("第二行为，true!")
	}*/
	/*//位运算符
	var a uint = 60
	var b uint = 13
	var c uint
	c = a & b
	fmt.Printf("a与b为%d\n", c)
	c = a | b
	fmt.Printf("a与b为%d\n", c)
	c = a ^ b
	fmt.Printf("a与b为%d\n", c)
	c = a << b
	fmt.Printf("a与b为%d\n", c)
	c = a >> b
	fmt.Printf("a与b为%d\n", c)
	*/
	//赋值运算符
	a := 60
	var c int = 20
	c += a
	fmt.Printf("+= 运算符实例，c 值为 = %d\n", c) //c=c+a
	c -= a
	fmt.Printf("-= 运算符实例，c 值为 = %d\n", c) //c=c-a
	c *= a
	fmt.Printf("*= 运算符实例，c 值为 = %d\n", c) //c=c*a
	c /= a
	fmt.Printf("/= 运算符实例，c 值为 = %d\n", c) //c=c/a
	c = 400
	c <<= 2
	fmt.Printf("<<= 运算符实例，c 值为 = %d\n", c) //c=c<<2
	c = 200
	c >>= 2
	fmt.Printf(">>= 运算符实例，c 值为 = %d\n", c) //c=c>>2
	c &= 2
	fmt.Printf("&= 运算符实例，c 值为 = %d\n", c) //c=c&2
	c ^= 2
	fmt.Printf("^= 运算符实例，c 值为 = %d\n", c) //c=c^2
	c |= 2
	fmt.Printf("|= 运算符实例，c 值为 = %d\n", c) //c=c|2

}
