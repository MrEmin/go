package main

import "fmt"

func main() {

	/* myArr := [3]int{1, 2, 3}
	myArr2 := [...]int{1, 2, 3, 4}
	fmt.Println(myArr)
	fmt.Println(myArr2)
	fmt.Println(len(myArr))
	fmt.Println(len(myArr2))

	// Array oluşturulduğunda eleman sayısı bilinirken, slice'da eleman sayısı bilinmemektedir.

	var myArr3 [5]int
	fmt.Println(myArr3)

	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	fmt.Println(len(mySlc)) */

	/* var myArr [4]int
	fmt.Println(myArr)
	myArr[0] = 5
	fmt.Println(myArr) */

	// go kodları copile ederken array'ler daha güvenlidir. Çünkü arraylerin sınırları bellidir ve daha hızlı compile edilir.
	// bir arrayin kopyasını alırken arrayin kendisini kopyalarız yani array'in değerini kopyalarız ancak slice'da böyle değildir.
	// slice'larda referansların kopyasını alırız.

	/* var mySlc []int
	mySlc = make([]int, 4)
	fmt.Println(mySlc)
	mySlc[0] = 10
	fmt.Println(mySlc) */

	/* myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)

	myArr2[0] = 100
	fmt.Println(myArr2)
	fmt.Println(myArr) */

	// Array'ler Pass By Value olarak tanımlanır.
	// Slice'lar ise Pass By Reference olarak tanımlanır.
	// Slice'lar array'ler üzerine oluşturulmuş olan bir yapıdır.

	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)

	mySlc2 := mySlc
	fmt.Println(mySlc2)

	mySlc2[0] = 33
	fmt.Println(mySlc2)
	fmt.Println(mySlc)
}
