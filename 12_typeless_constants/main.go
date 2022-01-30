package main

import "fmt"

func main() {

	/* const x = 5 // typeless

	// go x'e varsayılan bir veri tipi atıyor. int
	fmt.Printf("%T, %v\n", x, x) */

	/* const x = 3 // sabit typeless olduğu için varsayılan olarak int ataması yapar.
	var y int16 = 12

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y) // işlemi hatasız yapabilmek için typeless sabiti type conversion ile int -> int 16 yapar.
	fmt.Printf("%T, %v\n", x, x) */

	/* const x = int16(5.2 + 4.8)
	fmt.Printf("%T, %v\n", x, x) */

	/* const x = 3
	const y = 5.6

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y) */

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)
}
