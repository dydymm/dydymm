package main

import (
	"errors"
	"fmt"
)

//Go 语言习惯使用一个独立、明确的返回值来传递错误信息。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。Go 语言的处理方式能清楚的知道哪个函数返回了错误。

//错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 50 {
		return -1, errors.New("can't work in 50") //errors.New 使用给定的错误信息构造一个基本的 error 值。
	}
	return arg + 3, nil //返回错误值为 nil 代表没有错误。
}

//还可以通过实现 Error() 方法来自定义 error 类型。
type argError struct {
	arg  int
	prob string
}

//我们使用 &argError 语法来建立一个新的结构体， 并提供了 arg 和 prob 两个字段的值。
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 50 {
		return -1, &argError{arg, "can't work at there"}
	}
	return arg + 9, nil
}

func main() {

	//下面的两个循环测试了每一个会返回错误的函数。 注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。
	for _, i := range []int{3, 50} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 work :", r)
		}
	}
	for _, i := range []int{7, 50} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 work :", r)
		}
	}

	_, e := f2(50)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
