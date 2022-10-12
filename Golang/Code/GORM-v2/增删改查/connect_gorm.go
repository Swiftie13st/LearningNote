package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var GLOBAL_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:abc123@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 连接池最多容纳的连接数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接池中连接可复用的最大时间

	GLOBAL_DB = db

	TestUserCreate()
	//CreatedTest()
	//TestSelect()
	//TestUpdate()
	//TestDelete()
	TestSQL()
}
