package main

import "fmt"

func main() {

	/* for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	} */

	// for <init statement>; <condition>; <post statement> { ----- }

	/* i := 0

	for ; i <= 10; i += 5 { // baştaki ;'ü i'nin tanımlaması gibi düşün.
		fmt.Println(i)
	}

	fmt.Println(i) */

	/* for { // Infinite Loop
		fmt.Println("Hello World")
	} */

	/* for i := 0; true; i += 5 {
		fmt.Println(i)
	} */

	/* for i := 0; false; i += 5 {
		fmt.Println(i)
	} */

	/* i := 10

	for i >= 0 { // tek ifade olduğu için başa ; koymaya gerek yok.
		fmt.Println(i)
		i--
	} */

	/* for i := 0; i <= 10; i++ {

		if i%3 == 0 {
			continue // bu koşulu sağladığında aşağıya devam etmeden for döngüsünün başına gider.
		}

		fmt.Println(i)
	} */

	for i := 0; i <= 10; i++ {

		if i == 3 {
			break // break --> döngüden çık.
		}

		fmt.Println(i)
	}

}
