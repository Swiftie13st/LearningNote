package V1

// 导入mysql驱动
import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 定义模型
type User struct {
	// 内嵌gorm.Model
	gorm.Model
	Name     string
	Age      sql.NullInt64 // 零值类型
	Birthday *time.Time
	// 建立唯一索引
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段，不会存在数据库中
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 唯一指定表名
func (Animal) TableName() string {
	return "Animal_new"
}

func main() {

	// 关于默认表名的修改函数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:abc123@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	fmt.Println("connect mysql success")

	// 禁用表名的负数形式【默认会在表名后面加s】
	db.SingularTable(false)

	// 自动迁移【把结构体和数据表进行对应】
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	// 创建记录
	//u1 := User{
	//	Name: "Bruce",
	//}
	//
	//db.Create(&u1)
	//
	//// 查询数据
	//var u User
	//db.First(&u)
	//fmt.Printf("u: %#v \n", u)
	//
	//// 更新
	//db.Model(&u).Update("hobby", "双色球")
	//
	//// 删除
	//db.Delete(&u)

}
