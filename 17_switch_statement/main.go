package main

import (
	"fmt"
)

func main() {

	// Switch
	// Go'da varsayılan olarak case'lerin altında break; vardır.
	// Go'da switch yapısı, if - else yapısının daha sade halidir.

	grade := 4

	switch grade {
	case 5: // if grade == 5 { fmt.Println("Pekiyi") }
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("İyi")
		/* y := 100
		fmt.Println(y) */

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Geçer")

	case 1:
		fmt.Println("Başarısız")

	default: // Default'u istediğimiz alana yazabiliriz.
		fmt.Println("Geçersiz Not")
	}

	switch {
	case false:
		fmt.Println("Bu yazdığımız konsolda görünmez.")

	case true:
		fmt.Println("Bu yazdığımız konsolda görünecek.")
	}

	/* 	if grade == 5 {
	   		fmt.Println("Pekiyi")
	   	} else if grade == 4 {
	   		fmt.Println("İyi")
	   	} else if grade == 3 {
	   		fmt.Println("Orta")
	   	} else if grade == 2 {
	   		fmt.Println("Geçer")
	   	} else if grade == 1 {
	   		fmt.Println("Başarısız")
	   	} else {
	   		fmt.Println("Geçersiz Not")
	   	} */

}
