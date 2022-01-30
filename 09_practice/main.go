// 1 : int x, float64 y type conversion sample

/* package main

import "fmt"

func main() {

	x := 75
	var y float64
	y = float64(x) // type(value)

	fmt.Println(y)

} */

// 2 : multiple assing sample x, y = y, x

/* package main

import "fmt"

func main() {

	x := 5
	y := 10

	fmt.Println("X:", x, "Y:", y)

	x, y = y, x // x= y, y = x
	fmt.Println("X:", x, "Y:", y)

} */

// 3 : non English variable names

/* package main

import "fmt"

func main() {
	//yaş := 22
	//fmt.Println(yaş)

	名稱 := "Emin"
	年齡 := 22

	fmt.Println("Name:", 名稱, "Age:", 年齡)
} */

// 4 : shadowing kavramı? gölgeleme

/* package main

import "fmt"

func main() {
	x := 5

	if true {
		x := 10 // yeni bir x değişkeni oluşturulur ve sadece bu bloğun içinde çalışır.
		x++
		fmt.Println(x) // 11
	}

	fmt.Println(x) // 5
} */

// 5 : 40 as a string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 65
	s := string(x)

	fmt.Printf("%v, %T\n", x, x) // Çıktı => 65, int
	fmt.Printf("%v, %T\n", s, s) // ASCII'de 65. karakter A harfidir. Çıktı => A, string

	y := strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y) // Çıktı => 65, string
}
