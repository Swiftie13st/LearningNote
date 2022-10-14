package main

type TMG struct {
	ID   uint
	Name string
}

func Transaction() {
	GLOBAL_DB.AutoMigrate(&TMG{})

	//flag := false
	//GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
	//	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	//	tx.Create(&TMG{Name: "T1"})
	//	tx.Create(&TMG{Name: "T2"})
	//	tx.Create(&TMG{Name: "T3"})
	//	if flag {
	//		// 返回 nil 提交事务
	//		return nil
	//	} else {
	//		return errors.New("ERROR")
	//	}
	//
	//})

	tx := GLOBAL_DB.Begin()
	tx.Create(&TMG{Name: "T1"})
	tx.Create(&TMG{Name: "T2"})
	tx.SavePoint("sp")
	tx.Create(&TMG{Name: "T3"})
	tx.RollbackTo("sp") // 不会创建T3
	tx.Commit()

}
