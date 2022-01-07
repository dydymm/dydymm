package main
import (
	"fmt"
)
func main()  {

	//break
	for i := 1; i <= 10; i++{
		if i==6 {
			break
		}
		fmt.Println("i=",i)
	}
	fmt.Printf("\n")
	//continue
	for i := 1; i <= 10; i++{
		if i==6 {
			continue
		}
		fmt.Println("i=",i)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	//goto
	for i := 1; i <= 10; i++{
		if i==6 {
			goto END
		}
		fmt.Println("i=",i)
	}
		END :
		fmt.Println("END")


	//switch
	r:=1
	fmt.Printf("当r=%d时:\n",r)

	switch r {
		case 1://默认每个 case 带有 break
			fmt.Println("输出 1 r=",r)
			fallthrough//fallthrough 不跳出，并执行下一个 case
		case 2:
			fmt.Println("输出 2 r=",r)
		case 3,4:
			fmt.Println("输出3,4 r=",r)
		default:
			fmt.Println("输入错误，输入为:",r)
	}
}