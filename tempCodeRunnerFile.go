	var num = 10
	if num == 10 {
		fmt.Println("hello == 10")
	} else if num > 10 {
		fmt.Println("hello > 10")
	} else {
		fmt.Println("hello < 10")
	}

	if num2 := 10; num2 >= 10 {
		fmt.Println("hello >=10")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i+1)
	}