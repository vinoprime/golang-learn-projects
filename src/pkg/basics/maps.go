package basics

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(servers map[string]int) {
	fmt.Println("\n There are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println(stats[Online], "server are online")
	fmt.Println(stats[Offline], "server are Offline")
	fmt.Println(stats[Maintenance], "server are Maintenance")
	fmt.Println(stats[Retired], "server are Retired")
	fmt.Println(stats[Online], "server are online")
}

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

	fmt.Println("There are total items ", totalitems)

	// --- Chalenge

	servers := []string{"darkstart", "aiur", "omicron", "last"}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["darkstart"] = Retired
	serverStatus["aiur"] = Offline
	printServerStatus(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	printServerStatus(serverStatus)

}
