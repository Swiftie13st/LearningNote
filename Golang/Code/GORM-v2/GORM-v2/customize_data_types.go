package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// 自定义数据类型

type CInfo struct {
	Name string
	Age  int
}

// Value 实现 driver.Valuer 接口
func (c CInfo) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(str), nil
}

// Scan 实现 sql.Scanner 接口
func (c *CInfo) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	json.Unmarshal(str, c)
	return nil
}

type Args []string

// Value 实现 driver.Valuer 接口
func (a Args) Value() (driver.Value, error) {
	if len(a) > 0 {
		var str string = a[0]
		for _, i := range a[1:] {
			str += "," + i
		}
		return str, nil
	} else {
		return "", nil
	}

}

// Scan 实现 sql.Scanner 接口
func (a *Args) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型无法解析")
	}
	*a = strings.Split(string(str), ",")

	return nil
}

type CUser struct {
	ID   uint
	Info CInfo `gorm:"type:text"`
	Args Args
}

func Customize() {
	GLOBAL_DB.AutoMigrate(&CUser{})

	GLOBAL_DB.Create(&CUser{
		Info: CInfo{
			Name: "Bruce",
			Age:  18,
		},
		Args: Args{
			"1", "2", "3",
		},
	})

	var u CUser
	GLOBAL_DB.Last(&u)
	fmt.Println(u)
}
