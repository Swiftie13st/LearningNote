package main

import (
	"fmt"
	"log"
	"os"
)

func main1() {

	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err : %v \n", err)
		return
	}
	// 设置log的输出路径
	log.SetOutput(fileObj)

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
