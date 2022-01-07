package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { //area 是一个拥有 *rect 类型接收器(receiver)的方法。
	return r.width * r.height
}

func (r rect) perim() int { //可以为值类型或者指针类型的接收者定义方法。 这是一个值类型接收者的例子。
	return r.height*2 + r.width*2
}

func main() {
	r := rect{width: 20, height: 13}
	//调用为结构体定义的两个方法。
	println("area:", r.area())
	println("perim", r.perim())

	rp := &rect{width: 10, height: 2} //调用方法时，想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。
	fmt.Println("area:", rp.area()+1)
	fmt.Println("perim", rp.perim())
}
