package basics

import "fmt"

func Range() {

	slice := []string{"Hello", "world", "#"}

	for idx, element := range slice {
		fmt.Println(idx, element)

		for _, ch := range element {
			fmt.Printf("   %q\n", ch)
		}
	}
}
