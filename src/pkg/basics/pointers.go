package basics

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// * no need here because we are using struct
	counter.hits += 1
}

func incrementPrimite(counter *int) {
	*counter += 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool
type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out..")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}

}

func Pointers() {

	fmt.Println("Hello Pointer")

	// Syntax
	// value :=10
	// var valPtr *int
	// valPt = &value

	// value :=10
	// valptr := &value

	y := 5

	x := Counter{
		hits: 10,
	}

	increment(&x)
	fmt.Println(x)
	incrementPrimite(&y)
	fmt.Println(y)

	counter := Counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "hi", &counter)
	fmt.Println(hello, world)

	pharse := []string{hello, world}
	fmt.Println(pharse)

	replace(&pharse[1], "Go!", &counter)
	fmt.Println(pharse)

	/*Create a program that can activate and deactivate security tags on products*/
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	cloth := Item{"Cloth", Active}
	karch := Item{"Karch", Active}

	items := []Item{shirt, pants, purse, cloth, karch}
	fmt.Println(items)

	deactivate(&items[0].tag)
	fmt.Println("Items 0 deactivated", items)

	checkout(items)
	fmt.Println("Items", items)

}
