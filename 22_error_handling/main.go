/* package main

import (
	"errors" // errors --> package, New --> method
	"fmt"
)

func main() {
	// explicit tanımlama
	result, err := evenNum(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı:", result)
	}
}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("Hata: Çift sayı girmediniz.")
	}
	return num, nil // nil -> başlangıç değeri olmayan ifadelere verilen değerdir. Javadaki null'a eşittir.
} */

/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sRoot(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Go'yu öğrenmek istiyorsan veri tipini bilmek zorundasın.
// error interface olan bir veri tipidir.

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz.")
	}
	return math.Sqrt(num), nil
} */

/* package main

import (
	"fmt"
	"math"
)

// hata programını yazmazsak hatanın ne olduğunu bilemeyiz.

func main() {
	result := sRoot(-16)
	fmt.Println(result)
}

func sRoot(num float64) float64 {
	return math.Sqrt(num)
} */

// Go hatanın açık bir şekilde tanımlanıp, yakalanmasını ister.
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız:", file)
	}
}
