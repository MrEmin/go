/* package main

import "fmt"

func main() {

	// Fonksiyonu çalıştırırken kullandığımız değişkenlerede argüman deriz
	merhaba("Arin", 6) // Argument Function Run

}

func merhaba(name string, age int) { // Parameter Function Writer
	fmt.Printf("Adım %s, yaşım %d", name, age)
} */

/* package main

import "fmt"

func main() {
	fmt.Println(result(50))
}

// son yazdığımız return'den sonra ffarklı bir statement yazamayız.
func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"
} */

/* package main

import "fmt"

func main() {
	merhaba("Arin", 6)

	name := "Elis"
	age := 4
	fmt.Printf("Adım %s, yaşım %d\n", name, age)

}

// Bir fonksiyonda kullandığımız değişkenlerin scope'u (kapsam) o fonksiyon içerisindedir.
func merhaba(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d\n", name, age)
} */

// Fonksiyonları biz yazarız. Metotlar ise programın içindeki hazır fonksiyonlarıdır.
/* package main

import (
	"fmt"
	"time"
)

func main() {

	//var x int = 10

	var moment time.Time = time.Now() // Now --> method
	fmt.Println(moment)

} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
} */

// multiple return
package main

import "fmt"

func main() {
	bölüm, kalan := bölme(104, 5)
	fmt.Printf("Bölüm: %v, Kalan: %v", bölüm, kalan)
}

// 104 / 5 ===> 20 - 4

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}
