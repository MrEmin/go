package main

import "fmt"

func main() {

	// Neden fonksiyon kullanırız? Biz fonksiyonlar yardımıyla programımızı daha modüler hale getiririz. Daha okunabilir bir kod olur.

	/* var x, y, sum int

	x = 5
	y = 10
	sum = x + y
	fmt.Printf("%d ve %d sayısının toplamı %d\n", x, y, sum)

	x = 7
	y = 11
	sum = x + y
	fmt.Printf("%d ve %d sayısının toplamı %d\n", x, y, sum) */

	// Modular Programming
	// Readable code
	// From Complex To Simple

	fmt.Println(sum(5, 7))

	// func funcName(params) returnType { code }

	merhaba()
}

func sum(x, y int) int {
	return x + y
}

func merhaba() {
	fmt.Println("Merhaba Dünya")
}
