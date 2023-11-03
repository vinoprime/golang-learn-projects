package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func Dmain() {

	fmt.Println("Main")

	rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 12
	rolls := 1

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, "Die #", d, ":", rolled)
		}

		fmt.Println("Toltal rolled:", sum)
	}
}
