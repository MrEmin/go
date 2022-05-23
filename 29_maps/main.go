package main

import "fmt"

// map ifadesi key ve value değerlerinden oluşur.

func main() {
	/* myMap := map[string]int{
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}

	fmt.Println(myMap)
	// fmt.Println(myMap[0]) map'te index numarası yoktur. Key ile yazdırma işlemi yapılır.
	fmt.Println(myMap["Ahmet"], myMap["Selim"])
	fmt.Println(myMap["Ahmet Mehmet"]) */

	/* myMap := map[string]int{ // LITERAL
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Aslı"]) */ // olmayan bir key'e baktığımızda bize 0 döner.

	// key - value aynı veri tipinde olmasına gerek yok.
	// key ve value kendi içlerindeki veri tiplerinin aynı olması gerekir.

	/* 	isMarried := map[string]bool{
	   		"Ahmet":  true,
	   		"Ayşe":   false,
	   		"Mahmut": false,
	   	}

	   	fmt.Println(isMarried)

	   	// map'te her bir eleman ayrı ayrı değerlendirilir. */

	/* myMap := make(map[string]int)
	fmt.Println(myMap)
	fmt.Println(myMap["Test"]) */

	/* studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	} */

	/* fmt.Println(studentGrades)
	fmt.Println(studentGrades["Arin"])
	fmt.Println(studentGrades["Çınar"])
	// fmt.Println(studentGrades["Elis"])

	value, ok := studentGrades["Elis"] // value, değerini. ok ise map'te olup olmadığını söyler.
	fmt.Println(value, ok) */

	/* _, ok := studentGrades["Çınar"] // bu key'e sahip elemanın map'te olup olmadığını kontrol etmek için ok'u kontrol ederiz. Value'ye gerek yoktur. Buna comma ok control denir.
	fmt.Println(ok) */

	/* studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	} */

	/* fmt.Println(studentGrades)
	studentGrades["Mahmut"] = 55 // map'e bu şekilde yeni eleman ekleyebiliriz.
	fmt.Println(studentGrades)

	delete(studentGrades, "Selim") // map'e ait olan bir elemanı silmek için için gömülü delete metodu kullanılır.
	fmt.Println(studentGrades)

	fmt.Println(len(studentGrades)) */

	// MAP'ler de slice'lar gibi değeri değil referansları paylaşır. Pass By Reference

	studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}

	/* sg := studentGrades // maps --> pass by reference
	fmt.Println(studentGrades)
	fmt.Println(sg)
	fmt.Println("------------------------------------------")
	delete(sg, "Arin")
	fmt.Println(studentGrades)
	fmt.Println(sg) */

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
