package main

import "fmt"

func main() {

	/* x := 10

	// x if yapısını içersinde tanımlı olduğu için sadece if kapsamında çalışır.

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	fmt.Println(x) */

	// shadowing
	if x := -25; x < 0 { // tek satırda 2 statement yazabilmek için ikisinin arasına ; koymamız gerekiyor.
		fmt.Println(x, "negatif sayıdır") //Normalde go derleyicisi ;'ü arka planda kendisi koyuyor.
		fmt.Println("Benim adım Emin")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır")
		} else {
			fmt.Println(x, "tek sayıdır")
		}
	}

}
