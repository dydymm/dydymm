package main
import (
	"fmt"
)

	//数组中元素个数自动设置数组长度，并且不可改变。切片通过 append 方法增加元素
	//切片还有容量（cap）的概念。因此切片其实还有一个指定长度和容量的初始化方式
	//切片还可以从一个数组中初始化.
	//数组传递的是值，而切片传递的是指针。传入的切片在函数中被改变时，函数外的切片也会同时改变。相同的情况，函数外的数组则不会发生任何变化。

func main()  {
	person := [3]string {"TOM","Alean","Broun"}
	fmt.Printf("len=%d,cap=%d,array=%v\n",len(person),cap(person),person)

	fmt.Println("\n")
	for k,v := range person{
		fmt.Printf("person[%d]:%s\n\n",k,v)
	}

	for i := range person{
		fmt.Printf("person[%d]:%s\n\n",i,person[i])
	}

	for i := 0; i < len(person); i++ {
	fmt.Printf("person[%d]:%s\n",i,person[i])	
	}

	//使用空白符
	for _,name := range person{
		fmt.Println("name:",name)
	}

}