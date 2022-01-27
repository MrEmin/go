// https://go.dev/ref/spec#Types

package main

import (
	"fmt"
)

func main() {

	//var name string = "Emin"
	var name = "Emin"
	var age int16 = 22
	var isMarried bool = false
	var weight float32 = 80.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	// data types
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
}
