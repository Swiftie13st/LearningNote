# JWT

JWT是JSON Web Token的缩写，是一种跨域认证的解决方案。
[jwt.io](https://jwt.io/)

## why
在之前的一些web项目中，我们通常使用的是Cookie-Session模式实现用户认证。相关流程大致如下：

1. 用户在浏览器端填写用户名和密码，并发送给服务端
2. 服务端对用户名和密码校验通过后会生成一份保存当前用户相关信息的session数据和一个与之对应的标识（通常称为session_id）
3. 服务端返回响应时将上一步的session_id写入用户浏览器的Cookie
4. 后续用户来自该浏览器的每次请求都会自动携带包含session_id的Cookie
5. 服务端通过请求中的session_id就能找到之前保存的该用户那份session数据，从而获取该用户的相关信息。

这种方案依赖于客户端（浏览器）保存Cookie，并且需要在服务端存储用户的session数据。

在移动互联网时代，我们的用户可能使用浏览器也可能使用APP来访问我们的服务，我们的web应用可能是前后端分开部署在不同的端口，有时候我们还需要支持第三方登录，这下Cookie-Session的模式就有些力不从心了。

JWT就是一种基于Token的轻量级认证模式，服务端认证通过后，会生成一个JSON对象，经过签名后得到一个Token（令牌）再发回给用户，用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了。

## 格式

一个正确的JWT格式如下所示：
```
eyJhbGciOiJIUzI1NiIsInR5c.eyJ1c2VybmFtZaYjiJ9._eCVNYFYnMXwpgGX9Iu412EQSOFuEGl2c
```
我们看到一个JWT字符串由Header，Payload，Signature三个部分组成，中间使用逗号连接。

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202207041836106.png)

### Header

Header是一个JSON对象，由token类型和加密算法两个部分组成的，如：
```json
{
  "typ": "JWT",//默认为JWT
  "alg": "HS256"//支持多种加密算法
}
```

将上面的JSON对象使用Base64URL算法转换成字符串，即可得到JWT中的Header部分。

> 注意：JWT编码并不使用Base64,而Base64Url,这是因为Base64生成字符串里，可能会有+，/和=这三个URL中特殊的符号，而我们又可能将token放在URL上传递到服务器上(如test.com?token=xxx)， 而Base64URL算法，则是在Base64算法生成的字符串基础上，将=省略，将+替换成-，将/替换成_。

### Payload

JWT的Payload部分与Header一样，也是一个JSON对象，用来存放我们实际需要的数据，JWT标准提供了七个可选的字段，分别为：

| 标题 |  描述  | 
| :-: | :-: | 
| iss(issuer) | 签发者，其值为大小写敏感的字符串或Uri |
| sub(subject) | 主题,用于鉴别一个用户 | 
| exp(expiration time) | 过期时间 | 
| aud(audience) | 受众 | 
| iat(issued at) | 签发时间 | 
| nbf(not before) | 生效时间 | 
| jti(JWT ID) | 	编号 | 

除了标准的字段外，我们可以任意定义私有的字段以满足业务需求，如：
```json
{
    iss:"my",//标准字段
    jti:"test",//标准字段
    username:"aaa",//自定义字段
    "gender":"男",
    "avatar":"https://1.jpg"
}
```
将上面的JSON对象使用Base64URL算法转换成字符串，即可得到JWT中的Payload部分。

### Signature

Signature是JWT的签名，生成方式为：将Header与Payload进行Base64URL算法编码后，用逗号链接，再使用密钥(secretKey)和Header中指的加密方式进行加密，最终生成Signature。

## JWT的特点

1. 最好使用HTTPS协议,防止JWT被盗的可能。
2. 除了JWT签发时间到期外，没有其他办法让已经生成的JWT失效，除非服务器端换算法。
3. 在JWT不加密的情况下，JWT不应该存储敏感的信息，如果要存放敏感信息，最好再次加密。
4. JWT最好设置较短的过期时间，防止被盗用后一直有效，降低损失。
5. JWT的Payload也可以存储一些业务信息，这样可以减少数据库的查询。

## JWT的使用

服务器签发JWT后，发送给客户端，客户端如果是浏览器的话，可以将其存放在cookie或localStorage中，如果是APP的话，则可以存放在sqlite数据库中。
然后每一次接口请求时都带上JWT,而带上来给服务端的方式，也有很多种，比如query、cookie、header或者body，总之就是一切可以带上数据给服务器的方式都可以，但比较规范的做还是通过header Authorization上传,格式如下：
```
Authorization: Bearer <token>
```

# 在Go项目中使用JWT

使用 `github.com/golang-jwt/jwt`这个库来生成或解析JWT。

## 生成

使用`NewWithClaims()`方法生成Token对象，再通过Token对象的方法来生成JWT字符串，如：

```go
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	hmacSampleSecret := []byte("111") //密钥，不能泄露
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	//生成jwt字符串
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
}
```

也可使用New()方法生成Token对象，再生成JWT字符串,如：

```go
func main() {
	hmacSampleSecret := []byte("111")
	token := jwt.New(jwt.SigningMethodHS256)
	//通过New方法不能在创建的时候携带数据，因此可以通过给token.Claims赋值来定义数据
	token.Claims = jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
}
```
输出：
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.WhnRPdbXl6L6haLHSahb7ieyKLDppL764xAfiy3zlmk <nil>
```

上面的例子中，是通过`jwt.MapClaims`这个数据结构定义JWT中的Payload数据的，除了使用`jwt.MapClaims`外,我们也可以使用自定义的结构，不过该结构必须实现下面的接口：

```go
type Claims interface {
    Valid() error
}
```

下面是一个实现自定义数据结构的示例：

```go
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
}

func (c CustomerClaims) Valid() error {
	return nil
}

func main() {
	//密钥
	hmacSampleSecret := []byte("111")
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = CustomerClaims{
		Username: "张三",
		Gender:   "男",
		Avatar:   "https://1.jpg",
		Email:    "test@163.com",
	}
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
}
```

如果我们想在自定义结构中使用JWT标准中定义的字段，可以这样子：

```go
type CustomerClaims struct {
	*jwt.StandardClaims        //标准字段
	Username            string `json:"username"`
	Gender              string `json:"gender"`
	Avatar              string `json:"avatar"`
	Email               string `json:"email"`
}
```

## 解析

解析是生成反向操作，我们通过解析一个token来获取其中的Header,Payload,并通过Signature校验数据是否被窜改，下面是具体的实现：

```go
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
```

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202207041914737.png)

## 在Gin项目中使用JWT

通过上面的例子，再结合Gin框架，其实我们完全可以自己实现在Gin使用JWT的需求，但为了不重复造轮子，我们可以直接使用别人造好的轮子。
在Gin框架中，登录认证一般通过中间件来实现，而`github.com/appleboy/gin-jwt` [gin-jwt](https://github.com/appleboy/gin-jwt)这个库中已经集成`github.com/golang-jwt/jwt`的实现，并帮我们定义了对应的中间件和控制器。

下面是一个具体的例子

```go
package jwt

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//用于接受登录的用户名与密码
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

//jwt中payload的数据
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func JWT(r *gin.Engine) {

	// 定义一个Gin的中间件
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,   //刷新最大延长时间
		IdentityKey:      identityKey, //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
					// "FirstName": v.FirstName, // 
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: Authenticator, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			if v, ok := data.(*User); ok && (v.UserName == "admin" || v.UserName == "123") {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) { //错误时响应
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value.
		// This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//登录接口
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	{
		//退出登录
		auth.POST("/logout", authMiddleware.LogoutHandler)
		// 刷新token，延长token的有效期
		auth.POST("/refresh_token", authMiddleware.RefreshHandler)
		auth.Use(authMiddleware.MiddlewareFunc()) //应用中间件
		{
			auth.GET("/hello", helloHandler)
		}
	}

	// if err := http.ListenAndServe(":8005", r); err != nil {
	// 	log.Fatal(err)
	// }
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") || (userID == "123" && password == "123") {
		return &User{
			// 可以放在payload里
			UserName:  userID,
			LastName:  "Bo-Yi",
			FirstName: "Wu",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

//处理/hello路由的控制器
func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
		// "name":     user.(*User).FirstName + " " + user.(*User).LastName,
		"claims": claims,
		"user":   user,
	})
}
```
将服务器运行起来后，通过curl命令发起登录请求,如：

```bash
curl http://localhost:8005/login -d "username=admin&password=admin"
```

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202207041931534.png)

```json
{
    "code": 200,
    "expire": "2022-07-05T10:11:06+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY5ODcwNjYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1Njk4MzQ2Nn0.ZtQbwdnKSyy44CdCrOFNGTBj7sh6a6dY9ZYH7guyh3w"
}
```

### 请求需要token才能访问的接口：

#### 未带token访问时

```bash
curl http://localhost:8005/auth/hello
```

响应结果，如：
```
{"code":401,"message":"cookie token is empty"}
```

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202207041934884.png)

#### 带上token访问时

```bash
# 为了方便，先将上面获取的token设置为环境变量
export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY5MzgxMDksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1NjkzNDUwOX0.3APMFr_DzfHSL3GLRadKNgYzb-LKM67adXtxP-Ec2Zw

curl -H"Authorization: Bearer ${TOKEN}" http://localhost:8005/auth/hello
```

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202207041935833.png)

```json
// 不带jwt访问
{
    "code": 401,
    "message": "cookie token is empty"
}
// jwt过期
{
    "code": 401,
    "message": "Token is expired"
}
// 该jwt无权限访问
{
    "code": 403,
    "message": "you don't have permission to access this resource"
}
// 正确访问
{
    "text": "Hello World.",
    "userID": "admin",
    "userName": "admin"
}
```