package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s1 := "100"
	// i1, err := strconv.Atoi(s1)
	// if err != nil {
	// 	fmt.Println("can't convert to int")
	// } else {
	// 	fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	// }

	// b, err := strconv.ParseBool("true")
	// fmt.Println(b, err)

	// f, err := strconv.ParseFloat("3.1415", 64)
	// fmt.Println(f, err)

	// i, err := strconv.ParseInt("-2", 10, 64)
	// fmt.Println(i, err)

	// u, err := strconv.ParseUint("2", 10, 64)
	// fmt.Println(u, err)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s1, s2, s3, s4)
}
