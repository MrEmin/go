package main

import "fmt"

func main() {
	// if <boolean expression> {code} x%2 == 0 (false)

	/* x := 27

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	// ifadenin başına gelen "!" ifadeyi zıt haline çevirir. true veya false.

	if !false {
		fmt.Println("Mesaj Gösterilecek")
	}

	if 5 > 3 {
		fmt.Println("Mesaj Gösterilecek")
	} */

	/* x := 4

	// if'den sonra tek koşul varsa paranteze gerek yok.
	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	} */

	// if <boolean expression> { code } else { code }

	/* if true {
		fmt.Println("Mesaj Gösterilecek")
	} */

	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	// if <boolean expression> { code } else if <boolean expression> { code } else { code }
}
