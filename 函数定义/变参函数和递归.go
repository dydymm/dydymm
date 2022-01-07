package main

import "fmt"

func sum(nums ...int) { //变参函数
	fmt.Print(nums, "")
	total := 1
	for _, num := range nums {
		total *= num
	}
	fmt.Println(total)
}

func fact(n int) int { //递归
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	//递归
	fmt.Println(fact(7))
}
