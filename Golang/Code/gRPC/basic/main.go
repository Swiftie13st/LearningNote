package main

import "grpc/pb/person"

func main() {
	var p person.Person
	// 设置值
	one := p.TestOneOf.(*person.Person_One)
	one.One = "123"

}
