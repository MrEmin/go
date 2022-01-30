package main

/* var x = 10

var Y = 100 // Paket düzeyindeki değişkenler büyük harflerle isimlendirilir. */

func main() {
	// https://pkg.go.dev/fmt

	// (Print - Println) - Printf
	// Print ve Println raw string yazdırır. Printf ise formatlayıp yazdırır.

	/* fmt.Println("Merhaba")
	fmt.Print("Merhaba")
	fmt.Println("")
	fmt.Printf("Merhaba") */

	/* name := "Emin"

	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name)

	fmt.Print("Benim adım", name) // Çıktı => Benim adımEmin
	fmt.Println("")
	fmt.Println("Benim adım", name) // Çıktı => Benim adım Emin (değişkenden sonra bir boşluk bırakır ve yeni satıra geçer.)
	fmt.Printf("Benim adım", name)  // Çıktı => Benim adım%!(EXTRA string=Emin)

	fmt.Println()

	fmt.Printf("Benim adım %v %T", name, name) */ // %v => name değişkeninin değerini(value) göster demek. %T => değşkenin tipini yazdırır. %X ise 16'lık tabana göre büyük harflerle yazdırır.

	/* x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o", x, y, z) */

	// name, age := "Emin", 22

	// fmt.Print("Benim Adım ", name, " ve ben ", age, " yaşındayım.")
	// fmt.Println("Benim Adım", name, "ve ben", age, "yaşındayım.") // kelimele ve değişkenler arasına kendisi boşluk koyar.
	// fmt.Printf("Benim Adım %v ve ben %v yaşındayım.", name, age)

	// VISIBILITY - Görünülürlük

	// fmt.Println(x)

	/* 	// Go'da değişkenlerde camel case isimlendirme kullanılır.
	   	var coinType string
	   	var customerName string

	   	// kısaltmalar büyük harflerle yazılır.
	   	var URL
	   	var HTTP // yanında başka ifadeler olsa da kısaltmaları büyük harfle yazarız. "xyzHTTP"

	   	basit değişkenler i, j, k şeklinde tanımlanır. For döngülerinde olduğu gibi.
	*/

}
