package main

type Jiazi struct {
	ID          uint
	Name        string
	Xiaofengche []Xiaofengche `gorm:"many2many:jiazi_fenche;foreignKey:Name;joinForeignKey:jiazi;references:FCName;joinReferences:fengche;"`
}

type Xiaofengche struct {
	ID     uint
	FCName string
}

func Tags() {
	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Xiaofengche{})

	GLOBAL_DB.Create(&Jiazi{
		Name: "小夹子",
		Xiaofengche: []Xiaofengche{
			{FCName: "大风车"},
			{FCName: "大风车1"},
			{FCName: "大风车2"},
		},
	})
}
