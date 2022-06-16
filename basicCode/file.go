package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// // 读取文件 方法1
	// file, err := os.Open("./test.txt")
	// // 关闭文件流
	// defer file.Close()
	// if err != nil {
	// 	fmt.Println("打开文件出错")
	// }
	// // 读取文件里面的内容
	// var tempSlice = make([]byte, 1024)
	// var strSlice []byte
	// for {
	// 	n, err := file.Read(tempSlice)
	// 	if err == io.EOF {
	// 		fmt.Printf("读取完毕")
	// 		break
	// 	}
	// 	fmt.Printf("读取到了%v 个字节 \n", n)
	// 	strSlice := append(strSlice, tempSlice...)
	// 	fmt.Println(string(strSlice))
	// }

	// // 读取文件 方法2
	// file, err := os.Open("./test.txt")
	// // 关闭文件流
	// defer file.Close()
	// if err != nil {
	// 	fmt.Println("打开文件出错")
	// }
	// // 通过创建bufio来读取
	// reader := bufio.NewReader(file)
	// var fileStr string
	// var count int = 0
	// for {
	// 	// 相当于读取一行
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		// 读取完成的时候，也会有内容
	// 		fileStr += str
	// 		fmt.Println("读取结束", count)
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	count++
	// 	fileStr += str
	// }
	// fmt.Println(fileStr)

	// // 通过IOUtil读取
	// byteStr, _ := ioutil.ReadFile("./test.txt")
	// fmt.Println(string(byteStr))

	// 打开文件
	file, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 777)
	defer file.Close()
	str := "啦啦啦 \r\n"
	file.WriteString(str)

	// 通过bufio写入
	writer := bufio.NewWriter(file)
	// 先将数据写入缓存
	writer.WriteString("你好，我是通过writer写入的 \r\n")
	// 将缓存中的内容写入文件
	writer.Flush()

	// 第三种方式，通过ioutil
	str2 := "hello"
	ioutil.WriteFile("./test.txt", []byte(str2), 777)

	// 读取文件
	byteStr, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("读取文件出错")
		return
	}
	// 写入指定的文件
	ioutil.WriteFile("./test2.txt", byteStr, 777)
}
