package main


//"reflect"

func main() {
	/*数据类型，变量赋值
	a :=33.0
	b :=11.0
	fmt.Println("c=",a/b,reflect.TypeOf(a),reflect.TypeOf(b))
	*/
	/*格式化字符串
	stack :="post"
	ip :="192.168.0.1"
	port :=3389
	url :="http://%s:%d"
	target := fmt.Sprintf(url,ip,port)
	fmt.Println(target)
	fmt.Printf("send is:%s",stack)
	*/
	/*变量声明
	var aaa string
	var bbb int
	var ccc bool
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	*/
	/*//多行字符串

	str := `
	post
	ip:192.168.1.1
	port:6689
	`
	fmt.Println(str)
	*/

	/*//多变量声明
	var (
		aaa int
		bbb bool
	)
	fmt.Println(reflect.TypeOf(aaa), reflect.TypeOf(bbb))
	*/
	/*//常量
	const a, b, c = 1, "dym", false
	fmt.Println(a, b, c)
	*/
	/*//特殊常量
	const (
		a = iota
		b
		c
		d = "aaa"
		e
		f = iota
	)
	fmt.Println(a, b, c, d, e, f)
	*/
	/*//实例二
	const (
		x = 1 << iota
		y = 2 << iota
		z
		w
	)
	fmt.Println(x) //i=1：左移 0 位，不变仍为 1。
	fmt.Println(y) //j=3：左移 1 位，变为二进制 110，即 6。
	fmt.Println(z) //k=3：左移 2 位，变为二进制 1100，即 12。
	fmt.Println(w) //l=3：左移 3 位，变为二进制 11000，即 24。
	*/
}
