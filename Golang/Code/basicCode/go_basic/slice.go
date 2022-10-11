package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明切片，把长度去除就是切片
	var slice = []int{1, 2, 3}
	fmt.Println(slice)

	var slice2 []int
	fmt.Println(slice2 == nil)

	a := [5]int{55, 56, 57, 58, 59}
	// 获取数组所有值，返回的是一个切片
	b := a[:]
	// 从数组获取指定的切片
	c := a[1:4]
	// 获取 下标3之前的数据（不包括3）
	d := a[:3]
	// 获取下标3以后的数据（包括3）
	e := a[3:]

	fmt.Println(a, b, c, d, e)

	// 长度和容量
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s))

	ss := s[2:]
	fmt.Printf("长度%d 容量%d\n", len(ss), cap(ss))

	sss := s[2:4]
	fmt.Printf("长度%d 容量%d\n", len(sss), cap(sss))

	fmt.Println()
	var slices = make([]int, 4, 8)
	//[0 0 0 0]
	fmt.Println(slices)
	// 长度：4, 容量8
	fmt.Printf("长度：%d, 容量%d\n", len(slices), cap(slices))

	slices2 := []int{1, 2, 3, 4}
	slices2 = append(slices2, 5)
	fmt.Println(slices2)
	// 合并切片
	slices3 := []int{6, 7, 8}
	slices2 = append(slices2, slices3...)
	fmt.Println(slices2)
	// 输出结果  [1 2 3 4 5 6 7 8]

	// 需要复制的切片
	var slices4 = []int{1, 2, 3, 4}
	// 使用make函数创建一个切片
	var slices5 = make([]int, len(slices4), len(slices4))
	fmt.Println(slices5) // [0,0,0,0]
	// 拷贝切片的值
	copy(slices5, slices4)
	// 修改切片
	slices5[0] = 4
	fmt.Println(slices4)
	fmt.Println(slices5)

	// 删除切片中的值
	var slices6 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slices6)
	// 删除下标为1的值
	slices6 = append(slices6[:1], slices6[2:]...)
	fmt.Println(slices6)

	Sort_()

	var numSlice2 = []int{9, 8, 7, 6, 5, 4}
	sort.Ints(numSlice2)
	fmt.Println(numSlice2)

	// 逆序排列
	var numSlice4 = []int{9, 8, 4, 5, 1, 7}
	sort.Sort(sort.Reverse(sort.IntSlice(numSlice4)))
	fmt.Println(numSlice4)
}

func Sort_() {
	// 冒泡
	var numSlice = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(numSlice); i++ {
		flag := false
		for j := 0; j < len(numSlice)-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				var temp = numSlice[j+1]
				numSlice[j+1] = numSlice[j]
				numSlice[j] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(numSlice)

	// 选择排序
	var numSlice2 = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(numSlice2); i++ {
		for j := i + 1; j < len(numSlice2); j++ {
			if numSlice2[i] > numSlice2[j] {
				var temp = numSlice2[i]
				numSlice2[i] = numSlice2[j]
				numSlice2[j] = temp
			}
		}
	}
	fmt.Println(numSlice2)
}
