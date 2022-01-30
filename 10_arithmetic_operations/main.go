// addition + => 15 + 10 => 15, 10 operand, + operator
// substruction -
// product *
// division /
// remainder %

package main

import "fmt"

func main() {
	//x, y := 15, 10

	//fmt.Printf("%T, %v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("%T, %v\n", (x + y), (x + y))
	//fmt.Printf("%T, %v\n", (x - y), (x - y))
	//fmt.Printf("%T, %v\n", (x * y), (x * y))
	//fmt.Printf("%T, %v\n", (x / y), (x / y))
	//fmt.Printf("%T, %v\n", (9.0 / 3), (9.0 / 3))
	//fmt.Printf("%T, %v\n", (x % y), (x % y))

	//z := 5.0 / 2  //(2.5)
	//fmt.Printf("%T, %v\n", z, z)

	// Increment ++, Decrement --, POSTFIX
	// x++'nın println'de çalışmama sebebi: println bir statement'dir, aynı şekilde x++ veya x--'de statement'dir ve bir satırda sadece bir statement bulunabilir.

	x := 10
	/* fmt.Println(x)

	fmt.Println(x + 1)

	x++
	fmt.Println(x) */

	fmt.Println(x)

	x = x - 1
	fmt.Println(x)

	x--
	fmt.Println(x)
}

/*
//Python

print "hello" -> statement

5 * 5 -> #yields 25 -> expression

print 5 * 5 -> expression statement
*/
