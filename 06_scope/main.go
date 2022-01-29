package main

import "fmt"

// aynı değşken isimlerini farklı scope'larda kullanabiliriz fakat okunulurluk açısından iyi olmaz.

// paket değişkenlerinde short declaration yapamayız.
// paket düzeyindeki değişkenler programın çalışma süreci boyunca hafızada yer kaplar.
/* var packVar = "Package Scope"
var funcVar = "Func(Package) Scope" */

func main() {
	/* if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	} */

	var name = "Emin"

	// name değişkeni sadece atama yapar fakat surname ise hem declaration hemde atama yapar.
	name, surname := "Elis", "Software"
	fmt.Println(name, surname)

	// 2 adet aynı isimde değişken olduğu zaman, program kendi scope'u içindeki değişkeni yazdırır.
	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	/* 	fmt.Println(packVar)
	   	myFunc() */
}

/* func myFunc() {
	// paket düzeyindeki değişkeni yazdırır.
	fmt.Println(funcVar)
} */
