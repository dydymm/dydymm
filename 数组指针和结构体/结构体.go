package main

import (
	"fmt"
)

type student struct {
	name  string
	age   int
	hight float32
}

func main() {
	var teacher student
	teacher.name = "韩梅梅"
	teacher.age = 18
	teacher.hight = 100.8
	fmt.Println("韩梅梅:", teacher)

	var leader = student{name: "酸梅梅", age: 22}
	fmt.Println("酸梅梅:", leader)

	leader2 := student{name: "甜梅梅", hight: 179.223}
	fmt.Println("甜梅梅=", leader2)
	//匿名结构体
	comander := struct {
		name string
		age  int
		work string
	}{name: "亮晶晶", age: 49, work: "局长"}
	fmt.Println("亮晶晶:", comander)
}
