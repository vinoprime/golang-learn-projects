package basics

import "fmt"

func Variables() {
	var my_name = "Vio"
	fmt.Println(my_name)

	var name string = "VVVV"
	fmt.Println(name)

	user_name := "bbbb"
	fmt.Println(user_name)

	var sum int
	fmt.Println("The sum is ", sum)

	v1, v2 := 1, 5
	fmt.Println("split ", v1, v2)
}
