package calc

// 自定义包，最好和文件夹统一起来

// 公有变量
var age = 10

// 私有变量
var Name = "张三"

// 首字母大写，表示共有方法
func Add(x, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}
