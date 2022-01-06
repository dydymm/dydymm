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

}