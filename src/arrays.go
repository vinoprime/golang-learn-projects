package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checlCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is not cleaned")
		}
	}
}

type Product struct {
	price int
	name  string
}

func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost", cost)
}

func main() {
	fmt.Println("Array")

	// Declaration sytnatx
	// var array [3]int
	// array := [3]int{1,2,3}
	// array := [...]int{1,2,3,}
	// array :=[4]int{7,4,1}// 4th value would be 0 default value

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Storeroom"},
		{name: "Bathroom"},
	}

	checlCleanliness(rooms)

	fmt.Print("Preforming cleaning...")
	rooms[2].cleaned = true

	checlCleanliness(rooms)

	shoppingList := [...]Product{
		{1, "Banana"},
		{6, "Meat"},
		{8, "Stra"},
		{9, "Apple"},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{1, "Sweet Bread"}

	printStats(shoppingList)

}
