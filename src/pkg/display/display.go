package display

import "fmt"

/* Capitalized function name represents public function in Go*/
func Display(msg string) {
	print(msg)
}

/* Lower case function name represents private function in Go*/
func print(msg string) {
	fmt.Println(msg)
}
