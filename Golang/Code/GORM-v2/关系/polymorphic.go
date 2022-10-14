package main

type Jiazi struct {
	ID          uint
	Name        string
	Xiaofengche []Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:huhu"`
}

type Yujie struct {
	ID          uint
	Name        string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:Abaaba"`
}

type Xiaofengche struct {
	ID        uint
	Name      string
	OwnerType string
	OwnerID   uint
}

func Polymorphic() {
	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Yujie{}, &Xiaofengche{})

	GLOBAL_DB.Create(&Jiazi{
		Name: "夹子",
		Xiaofengche: []Xiaofengche{
			{Name: "小风车1"},
			{Name: "小风车2"},
		},
	})
	GLOBAL_DB.Create(&Yujie{
		Name: "御姐",
		Xiaofengche: Xiaofengche{
			Name: "大风车",
		},
	})
}
