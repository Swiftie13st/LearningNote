package main

import (
	"fmt"
	"goProject/calc"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	fmt.Printf("%v \n", calc.Add(2, 5))
	// 打印公有变量
	fmt.Println(calc.Name)

	_, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	var num1 float64 = 3.1
	var num2 float64 = 4.2
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)

	value := gjson.Get(json, "name.last")
	println(value.String())

}
