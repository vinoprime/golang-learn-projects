package basics

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calc(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs + rhs
	case Divide:
		return lhs / rhs
	default:
		panic("not implemented")
	}
}

func Iota() {
	fmt.Println("Hello iota")

	/* Constan like a variable , but unchanging */

	// const (
	// 	online = 0
	// 	offline = 1
	// 	Maintain = 2
	// )

	// const (
	// 	online = iota
	// 	offline
	// 	Maintain
	// )

	// const (
	// 	online   = iota + 3 // 0
	// 	_                   //1 (skip)
	// 	Maintain            //2
	// )

	// const (
	// 	online = iota +3
	// 	offline
	// 	Maintain
	// )

	add := Operation(Add)
	sub := Operation(Subtract)
	mul := Operation(Multiply)
	div := Operation(Divide)

	fmt.Println(add.calc(2, 2))
	fmt.Println(sub.calc(2, 2))
	fmt.Println(mul.calc(2, 2))
	fmt.Println(div.calc(2, 2))

}
