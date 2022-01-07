package main

import (
	"fmt"
	"math"
)

type jiheti interface { //定义几何体的接口
	area() float64
	perim() float64
}

//为rect和circle实现该接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//要实现一个接口，我们只需要实现接口中的方法。这里我们为 rect 实现了 jiheti 接口。
func (r rect) area() float64 { //area方法拥有一个名为r，类型为 rect的接收者。r是必要的！！
	return r.width * r.height
}
func (r rect) perim() float64 {
	return r.width*2 + r.height*2
}

//circle 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的 jiheti。
func measure(j jiheti) {
	fmt.Println(j)
	fmt.Println(j.area())
	fmt.Println(j.perim())
}

func main() {
	r := rect{width: 10, height: 7}
	c := circle{radius: 11}
	measure(r)
	measure(rect{width: 7, height: 2})
	measure(circle{})
	measure(c)
}
