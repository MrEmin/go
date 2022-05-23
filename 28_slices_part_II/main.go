package main

import "fmt"

func main() {
	/* underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:5] // [2 3 4] , 5. index dahil değil.

	mySlc2 := underArray[:6] //en baştan 6. indexe kadar al.
	fmt.Println(mySlc2)

	mySlc3 := underArray[3:] // 3. indexten sonuncuya kadar hepsini alır.
	fmt.Println(mySlc3)

	mySlc4 := underArray[:] // bütün elemanları alır
	fmt.Println(mySlc4)

	fmt.Println("-----------------------")

	fmt.Println(mySlc)
	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc4) */

	/* mySlc := []int{1, 2, 3} // mySliceUnderArray
	fmt.Println(mySlc)

	mySlc = append(mySlc, 4, 5)
	fmt.Println(mySlc)

	mySlc2 := append(mySlc, 4, 5) // mySlice2UnderArray oluşuyor. Burda slice'a yeni eleman ekliyoruz.
	fmt.Println(mySlc2) */

	// mySlc ve mySlc2 2 farklı underlying array'lere bakıyor.

	/* mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2) */

	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}

	mySlc = append(mySlc, mySlc2...)

	//3 nokta koyarak mySlc2'yi elemanlara ayırdık. Çünkü mySlc'ye sadece int elemenlar ekleyebiliriz. 3 nokta koymazsak slice türünde mySlc2'yi eklemeye çalışırsak hata alırız.

	// Go'yu öğrenceksek type'ları iyi öğrenmeliyiz.

	fmt.Println(mySlc) */
	/*
		mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(mySlc)
		// mySlc = mySlc[2:] //ilk n elemanı silmek için [n:]

		// mySlc = mySlc[:len(mySlc)-3] son n elemanı silmek için [:len(mySlc)-n]

		mySlc = mySlc[2 : len(mySlc)-3]

		fmt.Println(mySlc) */

	/* var myArr [4]int
	fmt.Println(myArr)

	var mySlc []int
	mySlc = make([]int, 4) // Make methodunu kullandığımızda array'deki gib 0 değerler oluşur. [0 0 0 0]
	fmt.Println(mySlc)

	var mySlc2 []bool
	mySlc2 = make([]bool, 2) // [false false]
	fmt.Println(mySlc2) */

	var mySlc3 []int          // herhangi bir değer atanmamış gözükür yani yok gözükür. Boş değildir.
	fmt.Printf("%#v", mySlc3) // []int(nil) , diğer dillerdeki null gibi

	fmt.Println()

	mySlc4 := make([]int, 0) // burda ise eleman olmadığı için boştur. Yok slice ile boş slice arasındaki fark budur.
	fmt.Printf("%#v", mySlc4)
}
