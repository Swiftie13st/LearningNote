// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

// func main() {
// 	// hmacSampleSecret := []byte("111") //密钥，不能泄露
// 	// //生成token对象
// 	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	// 	"foo": "bar",
// 	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
// 	// })
// 	// //生成jwt字符串
// 	// tokenString, err := token.SignedString(hmacSampleSecret)
// 	// fmt.Println(tokenString, err)

// 	hmacSampleSecret := []byte("111")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	//通过New方法不能在创建的时候携带数据，因此可以通过给token.Claims赋值来定义数据
// 	token.Claims = jwt.MapClaims{
// 		"foo": "bar",
// 		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
// 	}
// 	tokenString, err := token.SignedString(hmacSampleSecret)
// 	fmt.Println(tokenString, err)
// }
// package main

// import (
// 	"fmt"

// 	"github.com/golang-jwt/jwt"
// )

// type CustomerClaims struct {
// 	Username string `json:"username"`
// 	Gender   string `json:"gender"`
// 	Avatar   string `json:"avatar"`
// 	Email    string `json:"email"`
// }

// func (c CustomerClaims) Valid() error {
// 	return nil
// }

// func main() {
// 	//密钥
// 	hmacSampleSecret := []byte("111")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	token.Claims = CustomerClaims{
// 		Username: "张三",
// 		Gender:   "男",
// 		Avatar:   "https://1.jpg",
// 		Email:    "test@163.com",
// 	}
// 	tokenString, err := token.SignedString(hmacSampleSecret)
// 	fmt.Println(tokenString, err)
// }

package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type CustomerClaims struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func main() {
	var hmacSampleSecret = []byte("111")
	//前面例子生成的token
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IuW8oOS4iSIsImdlbmRlciI6IueUtyIsImF2YXRhciI6Imh0dHBzOi8vMS5qcGciLCJlbWFpbCI6InRlc3RAMTYzLmNvbSJ9.N-ZFu253w9PxQmuYbaCK-9DcvdfB74tH-B4-6tchjXc"

	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(t *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	claims := token.Claims.(*CustomerClaims)
	fmt.Println(claims)
}
