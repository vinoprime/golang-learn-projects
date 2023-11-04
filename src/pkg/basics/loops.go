package basics

import (
	"fmt"
)

func Loop() {

	fmt.Println("Main")

	for i := 1; i <= 50; i++ {
		divisible_by_3 := i%3 == 0
		divisible_by_5 := i%3 == 0

		if divisible_by_3 && divisible_by_5 {
			fmt.Printf("FizzBuzz")
		} else if divisible_by_3 {
			fmt.Printf("Fizz")
		} else if divisible_by_5 {
			fmt.Printf("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
