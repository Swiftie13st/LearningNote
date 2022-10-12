package main

import "fmt"

func TestSQL() {
	var users TestUser
	GLOBAL_DB.Raw("SELECT id, name, age FROM t_test_users WHERE name = ?", "Nobody").Scan(&users)
	fmt.Println(users)
}
