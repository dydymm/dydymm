package main
import (
	"encoding/json"
	"fmt"
)
func main()  {	
	/*映射
	//声明map
	//Map 集合是无序的 key-value 数据结构。
	var p1 map[int]string
	p1 = make(map[int]string)//使用make初始化
	p1[1]="TOM"
	fmt.Println("p1:",p1)

	var p2 map[int]string=map[int]string{}
	p2[1] ="TOm"
	fmt.Println("p2:",p2)

	var p3 map[int]string =make(map[int]string)
	p3[1]="Tom"
	fmt.Println("p3:",p3)
	
	p4 := map[int]string{}
	p4[1]="tom"
	fmt.Println("p4:",p4)

	p5:= make(map[int]string)
	p5[1]="toM"
	fmt.Println("p5:",p5)

	p6 := map[int]string{
		1:"tOM",
	}
	fmt.Println("p6:",p6)
	*/

	///*生成json
	res := make(map[string]interface{})//interface接口
	res["code"]=200
	res["msg"]="success"
	res["data"]=map[string]interface{}{
		"username":"TOM",
		"age":"30",
		"hobby":[]string{"读书","爬山"},
	}
	fmt.Println("map data:",res)
	
	//序列化
	jsons,errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:",errs)
	}
	fmt.Println("")
	fmt.Println("json data :",string(jsons))

	//反序列化
	res2 :=make(map[string]interface{})
	errs =json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:",errs)
	}
	fmt.Println("")
	fmt.Println("map data:",res2)

	//编辑和删除
	person := map[int]string{
		1:"tom",
		2:"alert",
		3:"walion",
	}
	fmt.Println("data:",person)

	delete(person,2)
	fmt.Println("data:",person)

	person[2]="jack"
	person[4]="kevin"
	fmt.Println("data:",person)

	//遍历map
	human := map[int]string{
		0:"张三",
		1:"李四",
		2:"王五",
	}
	fmt.Println("human =",human)
	human[6]="大朗"
	for k,v := range human{
		fmt.Println(k,v)	
	}
	//只遍历键
	for x := range human{
		fmt.Println(x)	
	}
	//只遍历值
	for _,y :=range human{
		fmt.Println(y)
	}

}