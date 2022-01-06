package main
import(
	"fmt"
)

func main(){
	/*
	//声明切片
	//切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。 len() 方法获取长度，cap() 可以测量切片最长可以达到多少。
	var sli_1 [] int //nil 切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int{}//空切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_2),cap(sli_2),sli_2)


	var sli_3 =[]int{1,2,3,4,5,6}
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_3),cap(sli_3),sli_3)


	sli_4 := []int{2,3,4,5,6,7,8}
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_4),cap(sli_4),sli_4)


	var sli_5 [] int =make([] int,5,8)//make来创建切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_5),cap(sli_5),sli_5)


	sli_6 := make([] int ,4,7)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli_6),cap(sli_6),sli_6)
	*/
	
	/*
	//截取切片
	sli :=[]int{5,6,7,8,9,0,1}
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli),cap(sli),sli)

	fmt.Println("sli[1] ==",sli[1])//从0开始第一位
	fmt.Println("sli[:] ==",sli[:])//整个数组
	fmt.Println("sli[1:] ==",sli[1:])//从第一位以后的数组
	fmt.Println("sli[:3] ==",sli[:3])//0-3位的数组
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli[0:3]),cap(sli[0:3]),sli[0:3])
	
	fmt.Println("sli[0:3:4] ==",sli[0:3:4])//4决定cap()大小
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli[0:3:4]),cap(sli[0:3:4]),sli[0:3:4])
	*/



	//追加切片
	sli:=[]int {4,5,6,7}//初始化数组
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli),cap(sli),sli)
//append 时，容量不够需要扩容时，cap 会翻倍。
	sli=append(sli,9)//追加切片
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli),cap(sli),sli)

	sli=append(sli,12,33,44)//追加多个元素
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli),cap(sli),sli)

	//删除切片
	//删除尾部两个元素
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli[:len(sli)-2]),cap(sli[:len(sli)-2]),sli[:len(sli)-2])
	//删除开头两个元素
fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli[2:]),cap(sli[2:]),sli[2:])
	//删除中间两个元素
	sli=append(sli[:3],sli[3+2:]...)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(sli),cap(sli),sli)
}
