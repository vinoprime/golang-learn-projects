package main

import "fmt"

const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func calc(v int) int {
	return v + 1
}

func weekDay(day int) bool {
	return day <= 4
}

func av(a, b, c int) float32 {
	return float32(a+b+c) / 3
}
func Omain() {

	q, w, e := 9, 7, 8

	if q > w {
		fmt.Println("a is higher")
	} else if q < w {
		fmt.Println("a is lower")
	} else {
		fmt.Println(e)
	}

	if av(q, w, e) > 7 {
		fmt.Println("Big")
	}

	fmt.Println()

	today, role := Tuesday, Manager

	if role == Admin || role == Manager {
		fmt.Println("Granted")
	} else if role == Contractor && !weekDay(today) {
		fmt.Println("Granted")
	} else if role == Member && !weekDay(today) {
		fmt.Println("Granted")
	} else if role == Guest {
		fmt.Println("Denied")
	}

	// Switch
	/*Switch*/

	x := 1
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}

	switch res := calc(5); {
	case res > 5:
		fmt.Println("5")
	default:
		fmt.Println("Default")
	}

	// Case list
	switch x {
	case 1, 2, 3:
		fmt.Println("1243")
	}

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	var j = 0
	for j < 5 {
		j++
		fmt.Println(j)
	}

	// Infinite Loop
	for {
		j++
		if j == 10 {
			fmt.Println("Finished")
			break
		}
	}

}
