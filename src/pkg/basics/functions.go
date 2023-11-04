package basics

import "fmt"

func double(x int) int {
	return x * x
}

func disp() {
	fmt.Println("Display")
}

func twoTwonumbers() (int, int) {
	return 1, 2
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func hi(name string) string {
	return "Hi" + name
}
func Functions() {
	s := double(2)
	disp()

	x, y := twoTwonumbers()

	greet("Vinoth")

	fmt.Println(s, x, y, hi("Vino"))
}
