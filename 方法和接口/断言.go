package main

import (
	"fmt"
)

//类型断言 提供了访问接口值底层具体值的方式。
func main() {
	var i interface{} = "hello"

	//语句断言接口值 i 保存了具体类型 string，并将其底层类型为 string 的值赋予变量 s。
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
