package main

import (
	"fmt"
)

/*func max(val1, val2 int) (int, int) { //值传递，不会影响实际参数
	var max int
	if val1 > val2 {
		max = val1
		return max, val1
	} else {
		max = val2
		return max, val2
	}
}
*/
/*
func swap(val1 *int, val2 *int) { //引用传递，传递变量的地址到函数，会影响到实际参数。
	var tmp int
	tmp = *val1
	*val1 = *val2
	*val2 = tmp
}
*/
//方法
//方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

type Circle struct { //定义结构体
	radius float64
}

func main() {
	/*函数定义部分
	a := 40
	b := 69
	//fmt.Println(max(a, b)) //值传递
	fmt.Println("变换前a=", a, "b=", b)
	//swap(&a, &b) //引用传递
	fmt.Println("变换后a=", a, "b=", b)*/

	//以下为方法部分
	var c Circle
	fmt.Println("边长为:", c.radius)
	c.radius = 10.00
	fmt.Println("边长为:", c.radius, "面积为:", c.getArea())
	c.changeRadius(20)
	fmt.Println("修改后，边长为:", c.radius)
	change(&c, 30)
	fmt.Println("传入指针的方式改变值,边长为:", c.radius)
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius
}
func (c *Circle) changeRadius(radius float64) { // 注意如果想要更改成功c的值，这里需要传指针
	c.radius = radius
}

//func (c Circle) changeRadius(radius float64)  {
//   c.radius = radius
//}
func change(c *Circle, radius float64) { // 引用类型要想改变值需要传指针
	c.radius = radius
}
