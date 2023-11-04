package basics

import "fmt"

type part string

func showLine(line []part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func Slices() {
	fmt.Println("Slice")

	/*
		Syntax and declaration

		mySlice := []int{1,2,3}
		item := mySlice[0]

		mySlice[startIdx : endIdx ]


		numbers :=[...]int{1,2,3,4}
		slice1 := numbers[:]  // [1,2,3,4]
		slice2 := numbers[1:] // [2,3,4]
		slice2 := slice2[:1] // [2]

		slice4 := numbers[:2] //[1,2]
		slice4 := numbers[1:3] //[2,3]

		num := []int{1,2,3}
		num = append(num,4,5,6)

		n1:[]int{1,2,3}
		n2:[]int{4,2,3}
		combined := append(n1.n2...)


		// Preallocate
		slice := make([]int,10)

		func iterate(slice []int) {
			for i:=0; i < len(slice); i++{

			}
		}

		small := []int{1}
		big := []int{1,2,3,4,5,6,7}
		iterate(small)
		iterate(big)


		// Multidimentional slice
		board := [][] string {
			[]string{"A","",""},
			[]string{"a","",""},
			{"_","_","_"}
		}

		board[0][0] = "X"
		board[2][2] = "0"

	*/

	/* Challenge:
	Create a function to print out the contents of the assembly parts
	*/

	assemblyLine := make([]part, 3)

	assemblyLine[0] = "Nut"
	assemblyLine[1] = "Bold"
	assemblyLine[2] = "Sheet"
	fmt.Println("3 Parts :")
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two parts")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced")
	showLine(assemblyLine)

}
