// 1 -) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.

package main

import "fmt"

/* func main() {
	cities := [5]string{"istanbul", "ankara", "izmir", "ordu", "bursa"}
	number := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}
	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])
	}
	fmt.Println()

} */

// 2 -) [0,1,2,3,4,5,6,7,8] slice'dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluşturunuz.

func main() {
	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	mySlc := number[:4]
	fmt.Println(mySlc)

	mySlc1 := number[4:7]
	fmt.Println(mySlc1)

	mySlc2 := number[6:]
	fmt.Println(mySlc2)

	mySlc3 := append(number[2:4], number[6:8]...)
	fmt.Println(mySlc3)
}

// 3 -) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.

// 4 -) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for rarne ile gösterin.
