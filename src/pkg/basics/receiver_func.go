package basics

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// Receiver function
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true

}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func ReceiverFunc() {
	fmt.Println("Hello Receiver")

	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println(lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)

	fmt.Println(lot)

	lot.vacateSpace(2)
	fmt.Println(lot)

}
