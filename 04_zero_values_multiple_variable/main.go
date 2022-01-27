package main

import (
	"fmt"
)

func main() {

	/* var (
		name      string = "Emin"
		age       int16  = 22
		isMarried bool   = false

		weight float32 = 80.5
		height int     = 180
	) */

	/* var name, age, isMarried, weight, height = "Emin", 22, false, 80.5, 180 */

	/* name, age, isMarried, weight, height := "Emin", 22, false, 80.5, 180 */

	var age bool

	fmt.Println(age) // string -> "", int || float -> 0, bool -> false - Zero Values
	/* fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height) */

}
