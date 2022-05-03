package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter a Number :")
	fmt.Scanf("%d", &num)
	// n := num / 2 // this is equal to bello check in if
	if n := num / 2; num%2 == 0 {
		fmt.Printf("%d is Even, %d times", num, n)
	} else {
		fmt.Printf("%d is Odd, %d times", num, n)
	}
}
