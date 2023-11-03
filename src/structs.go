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

}
