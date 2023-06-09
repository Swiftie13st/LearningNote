package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"goProject/tcp/server/proto"
)

// 用于接收请求的方法
func processConn(conn net.Conn) {
	// 与客户端通信
	var tmp [128]byte
	// 使用for循环监听消息
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(conn, string(tmp[:n]))
	}
}

func process_ori(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {

		// msg, err := proto.Decode(reader)
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	// // 本地端口启动服务
	// listen, err := net.Listen("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	fmt.Println("start server on failed ", err)
	// }

	// // for循环监听
	// for {
	// 	// 等待别人来建立连接
	// 	conn, err := listen.Accept()
	// 	if err != nil {
	// 		fmt.Println("accept failed, err: ", err)
	// 		return
	// 	}
	// 	go processConn(conn)
	// }

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
