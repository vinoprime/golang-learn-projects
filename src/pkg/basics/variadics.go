package basics

import "fmt"

// any number of variables we can gives, basically it is slice
func sum(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}
	return sum
}

func VariadicFunc() {
	fmt.Println("Hello Variadic")

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(all...)
	answerA := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(answer, answerA)
}
