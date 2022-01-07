package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 8
	var mean float32

	mean = float32(sum) * float32(count)
	fmt.Printf("mean的值为:%f\n", mean)

	var aa int64 = 44
	var bb int32
	//bb = aa报错cannot use aa (type int64) as type int32 in assignment

	bb = int32(aa) //正确
	fmt.Printf("bb=%d", bb)
}
