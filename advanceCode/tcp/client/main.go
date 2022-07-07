package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"goProject/tcp/server/proto"
)

func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// tcp client
func main() {
	// // 与server端建立连接
	// conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	fmt.Println("dial 127.0.0.1:20000 failed, err：", err)
	// 	return
	// }
	// // 发送数据
	// conn.Write([]byte(getInput()))

	// // 关闭流
	// defer conn.Close()

	// conn, err := net.Dial("tcp", "127.0.0.1:30000")
	// if err != nil {
	// 	fmt.Println("dial failed, err", err)
	// 	return
	// }
	// defer conn.Close()

	// // 连续发送20次的hello到服务器
	// for i := 0; i < 20; i++ {
	// 	msg := `Hello, Hello. How are you?`
	// 	conn.Write([]byte(msg))
	// }
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
