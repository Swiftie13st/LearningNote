package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var smp sync.Map

	smp.Store("k1", "v1")
	smp.Store("k2", "v2")
	smp.Store("k3", "v3")

	for i := 0; i < 10000; i++ {
		smp.Store("k"+strconv.Itoa(i), "v"+strconv.Itoa(i))
	}

	smp.Delete("k3")

	v, ok := smp.Load("k1")
	if !ok {
		fmt.Println("key not found")
	}

	fmt.Println(v.(string))

	smp.Range(func(k, v any) bool {
		fmt.Println(v.(string))
		return true
	})

}
