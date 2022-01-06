package main
import(
	"fmt"
)

func main(){
	//声明切片
	//切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。
	var sli_1 [] int //nil 切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int{}//空切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_2),cap(sli_2),sli_2)


	var sli_3 =[]int{1,2,3,4,5,6}
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_3),cap(sli_3),sli_3)


sli_4 := []int{2,3,4,5,6,7,8}
fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_4),cap(sli_4),sli_4)


var sli_5 [] int =make([] int, 5,8)
fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_5),cap(sli_5),sli_5)
}
