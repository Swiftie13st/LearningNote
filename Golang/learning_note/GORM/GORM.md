# GORM

## GORM 关系

- 一对一关系`Belongs To`， `Has One`
- 一对多关系`Has Many`
- 多对多关系`many to many`

## 在使用 GORM 进行数据库查询时，如何避免 N+1 查询的问题？

N+1 查询问题指的是在查询关联表时，如果使用了嵌套循环进行查询，就会产生大量的 SQL 查询。为了避免这个问题，可以使用 GORM 的 Preload 方法预先加载关联数据。

Preload主要作用是**在处理某个Model 时将关联的Model 中的信息预先读取出来**

### GORM 的 Preload 方法和 Joins 方法有什么区别？在什么情况下使用哪种方法更好？

Preload 方法是在查询时预加载关联数据，而 Joins 方法是通过 SQL JOIN 语句连接多个表查询数据。**Preload 方法适用于关联表较少、数据量不大的情况；而 Joins 方法适用于关联表较多、数据量较大的情况。**

## 事务处理 有哪些主要注意的点？

GORM 的事务管理使用 `Begin`、`Commit` 和 `Rollback` 方法实现。

-   首先启动事务时一定要做 **错误判断**。
-   建议在 **启动事务** 之后马上写 **defer方法**。
-   在 **defer方法** 内对 **err** 进行判断，如果 **全局** 中有 **err!=nil** 就回滚 (全局中 **err** 都为 **nil** 才能 **提交事务**)
-   在 **提交事务** 之后我们可以定义一个 **钩子函数 afterCommit**，来统一处理事务提交后的 **逻辑**。

```go
tx, err := g.DB().Begin()  // 启动事务
if err != nil {
   return errors.New("启动事务失败")
}

defer func() {
   if err != nil {
      tx.Rollback()  // 回滚
   } else {
      tx.Commit()    // 事务完成
      //定义钩子函数
      afterCommmit()
   }
}()
```

## 如何在 GORM 中使用原生 SQL 查询？

在 GORM 中，可以使用 `Raw` 方法来执行原生 SQL 查询。Raw 方法接受一个 SQL 查询字符串和可选的参数列表，并返回一个`*gorm.DB`对象，可以使用该对象进行进一步的查询操作。

## gorm有什么缺点？

**优点：**

-   提高开发效率

**缺点:**

-   牺牲执行性能【中间多了一个环节】
-   牺牲灵活性
-   弱化SQL能力

## 如何防止sql注入

### sql注入了解吗？

SQL 注入是一种专门针对SQL语句的攻击方式。通过把SQL命令插入到web表单提交、输入域名或者页面请求的查询字符串中，利用现有的程序，来非法获取后台的数据库中的信息。在web的测试中涉及到的会比较多些。

存在注入的原因是后台在编写程序时，没有对用户输入的数据做过滤。

1. 用户在某个输入框提交的参数是123
2. 浏览器提交的URL为: http://www.xxx.com/index.php?id=123
3. 服务器后台执行SQL语句：select * from table1 where id = 123此时是没有任何影响的。
4. 如果用户提交的参数是 123;drop table
5. 服务器后台执行SQL语句： select * from table1 where id =123 ; drop table
6. 相当于后台执行了两条SQL语句，查表，并且把table删除, 从而导致了SQL注入。

### 防止

在 Gorm 中，就为我们封装了 SQL 预编译技术，可以供我们使用。其中预编译的 SQL 语句merchant_id = ?和 SQL 查询的数据merchantId将被分开传输至 DB 后端进行处理。

```go
db = db.Where("merchant_id = ?", merchantId)
db = db.Where(fmt.Sprintf("merchant_id = %s", merchantId)) // 有风险
```

## 如何批量插入某个字段 

可以使用GORM的 DB 对象来执行原生 SQL 操作，然后结合 SQL 的 INSERT INTO 语句来实现批量插入数据。

要有效地插入大量记录，请将切片传递给 Create 方法。 GORM 将生成一条 SQL 语句来插入所有数据并回填主键值，钩子方法也会被调用。当记录可以分成多个批次时，它将开始一个事务。
您可以在使用 CreateInBatches 创建时指定批量大小

```go
var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}  
  
// batch size 100  
db.CreateInBatches(users, 100)
```

## 更新有几种方式 

1. `Save` 会保存所有的字段，即使字段是零值db.Save(&user)
2. `update` 使用 Update 更新单个列时，它需要有任何条件，否则会引发错误db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
3. 原生sql：GORM 允许使用 SQL 表达式更新列