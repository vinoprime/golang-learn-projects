package main

import "fmt"

type Sample struct {
	field string
	a, b  int
}

type Passenger struct {
	Name        string
	TickeNumber int
	Boarded     bool
}

type Bus struct {
	FrontSeat Passenger
}

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate //top left
	b Coordinate //bottom left
}

func width(rec Rectangle) int {
	return (rec.b.x - rec.a.x)
}

func length(rec Rectangle) int {
	return (rec.a.y - rec.b.y)

}

func area(r Rectangle) int {
	return length(r) * width(r)
}

func perimeter(r Rectangle) int {
	return (width(r) * 2) + (length(r) * 2)
}

func printInfo(r Rectangle) {
	fmt.Println("Area", area(r))
	fmt.Println("Perimeter", perimeter(r))
}

func Stmain() {

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

	// Anonymous
	var sample struct {
		f string
	}

	sample.f = "he"
	fmt.Println(sample)

	samp := struct {
		a int
	}{1}

	fmt.Println(samp)

	p1 := Passenger{"V", 1, false}
	fmt.Println(p1)

	var (
		bill = Passenger{Name: "V1", TickeNumber: 2}
		elen = Passenger{Name: "Ella", TickeNumber: 3}
	)

	fmt.Println(bill, elen)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TickeNumber = 4
	fmt.Println(heidi)

	p1.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	// Rectangle
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect)

	rect.a.y *= 2
	rect.b.x *= 2
	printInfo(rect)

}
