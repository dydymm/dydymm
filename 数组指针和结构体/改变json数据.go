package main
import (
	"encoding/json"
	"fmt"
)
type Result struct{
	Code int `json:"code"`
	Message string `json:"msg"`
	Name string `json:"name"`
}

func main(){
	var res Result
	res.Code =200
	res.Message ="success"
	res.Name = "dydymm"

	toJson(&res)
	setData(&res)
	toJson(&res)
}


func toJson (res *Result){//序列化
	jsons,errs := json.Marshal(res)
	if errs != nil{
		fmt.Println("json marshal error:",errs)
	}
	fmt.Println("json data:",string(jsons))
}

func setData (res *Result){//设置数据
	res.Code =404
	res.Message ="error"
	res.Name = "mmdydy"
}