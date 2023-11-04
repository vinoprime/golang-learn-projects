package basics

import "fmt"

func Maps() {

	/*  Syntax
	map[<key_type>]<value_type>
	map[string]int

	myMap := make(map[string]int)

	mymap := map[string]int{
		"key-1",1
		"key-2",12132
	}

	Operations:
	Inser => mymap["fav number"] = 5
	Read => favorite := myMap["fav number"]
	Delete => delete(mymap, "fav number")

	price, found := myMap["price"]
	if !found {
		// price not found
		return
	}


	*/

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	fmt.Println(shoppingList)
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"], "Eggs")

	beans, found := shoppingList["beans"]
	if !found {
		fmt.Println("not found")
	} else {
		fmt.Println(beans)
	}

	totalitems := 0
	for _, amount := range shoppingList {
		totalitems += amount
	}

	fmt.Println("There are ", totalitems)

}
