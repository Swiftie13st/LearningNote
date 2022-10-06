package main

import "fmt"

func main() {
	// // 向标准输出写入内容
	// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	// fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("打开文件出错，err:", err)
	// 	return
	// }
	// name := "沙河小王子"
	// // 向打开的文件句柄中写入内容
	// fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	// s1 := fmt.Sprint("沙河小王子")
	// name := "沙河小王子"
	// age := 18
	// s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	// s3 := fmt.Sprintln("沙河小王子")
	// fmt.Println(s1, s2, s3)

	// err := fmt.Errorf("这是一个错误")
	// fmt.Println(err)
	// e := errors.New("原始错误e")
	// w := fmt.Errorf("Wrap了一个错误:%w", e)
	// fmt.Println(w)

	// fmt.Printf("%v\n", 100)
	// fmt.Printf("%v\n", false)
	// o := struct{ name string }{"小王子"}
	// fmt.Printf("%v\n", o)
	// fmt.Printf("%#v\n", o)
	// fmt.Printf("%T\n", o)
	// fmt.Printf("100%%\n")

	// n := 12.34
	// fmt.Printf("%f\n", n)
	// fmt.Printf("%9f\n", n)
	// fmt.Printf("%.2f\n", n)
	// fmt.Printf("%9.2f\n", n)1
	// fmt.Printf("%9.f\n", n)\
	var (
		name    string
		age     int
		married bool
	)
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
