package main

import "fmt"

type Sample struct {
	field string
	a, b  int
}

func main() {

	data := Sample{
		field: "Hello",
		a:     1,
		b:     2,
	}

	d := Sample{}

	d.field = "HelloV"
	d.a = 100

	fmt.Println(d)
	fmt.Println(data)
	fmt.Println(data.field)

}
