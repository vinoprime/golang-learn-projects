package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num int
	// fmt.Printf("Enter a number : ")
	// fmt.Scanf("%d", &num)
	// // evenOrOdd(num)
	// fmt.Printf("%s", fizzBuzz(num))

	// switchCase()
	// forLoop()
	// whileLoop()
	// infiniteLoop()
	// array()
	slice()

}

func slice() {
	// empty slice
	var list []int
	list = append(list, 1, 2)
	printSliceInt(list)

	daysOfWeek := [...]string{
		"sun", "mon", "tue", "wed", "thu", "fri", "sat",
	}

	weekDays := daysOfWeek[1:6]
	weekDays[2] = "wednessday"
	printSlice(daysOfWeek[:])
	printSlice(weekDays)
}

func printSliceInt(slice []int) {
	fmt.Printf("%v\n", slice)
}

func printSlice(slice []string) {
	fmt.Printf("%v\n", slice)
}

func array() {
	var arr [5]int
	arr = [...]int{78, 56, 89, 23, 98}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%d: %d ", i, arr[i])
	// }

	// for i, n := range arr {
	// 	fmt.Printf("%d: %d\n", i, n)
	// }

	for _, n := range arr {
		fmt.Printf("%d\n", n)
	}
}

func infiniteLoop() {
	var temp, sum int
	var str string
	var err error
	for {
		fmt.Scanf("%s\n", &str)
		temp, err = strconv.Atoi(str) // with callback error handling
		// nil == null
		if err != nil {
			break
		}
		sum = sum + temp
	}
	fmt.Printf("%d", sum)
}

func whileLoop() {
	var temp, sum int
	i := 0
	for i < 5 {
		fmt.Scanf("%d\n", &temp)
		sum = sum + temp
		i++
	}
	fmt.Printf("%d", sum)
}

func forLoop() {
	var temp, sum int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d\n", &temp)
		sum = sum + temp
	}
	fmt.Printf("%d", sum)
}

func switchCase() {
	var c byte
	fmt.Printf("Say (y/n)")
	fmt.Scanf("%c", &c)
	// c = byte(unicode.ToLower(rune(c)))
	switch c {
	case 'y', 'Y':
		fmt.Printf("Thank you!")
	case 'n', 'N':
		fmt.Printf("Please...")
	default:
		fmt.Printf("Invlid input")
	}
}

func evenOrOdd(num int) {
	// n := num / 2 // this is equal to bello check in if
	if n := num / 2; num%2 == 0 {
		fmt.Printf("%d is Even, %d times", num, n)
	} else {
		fmt.Printf("%d is Odd, %d times", num, n)
	}
}

func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
