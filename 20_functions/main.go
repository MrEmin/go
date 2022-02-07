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

	/* fmt.Println(sum(5, 7))
	fmt.Println(sum(3, 5))
	fmt.Println(sum(2, 9))

	// func funcName(params) returnType { code }

	merhaba()
	merhaba() */

	// return vs print

	z := sum(5, 10)
	fmt.Println(z)

	sum2(6, 11)

	merhaba2("Muhammet Emin", 22)

	// Fonksiyon İsimlendirme
	// İlk karakter harf olacak.
	// camelCase isimlendirme kullanılır --> mySum
	// paket dışında kullanacaksak --> ilk harf büyük olur.
}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Merhaba Dünya")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
