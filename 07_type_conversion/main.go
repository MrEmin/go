package main

import "fmt"

func main() {
	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion -> type(value) => int(y) => 10.0 => 10

	fmt.Println(x + int(y))

	// Type Conversion ile yeni bir değer oluşturulur ancak değişkenin asıl veri tipini ve asıl değerini değiştirmiyoruz.
	fmt.Printf("%v %T\n", y, y) // 10 float64 */

	/* var x int8 = 10
	var y int16 = 10

	fmt.Println(x + int8(y)) */

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y) */

	// Type Conversion yaparken küçük olan veri tipini büyük olan veri tipine dönüştürmek daha sağlıklıdır, böylelikle veri kaybı olmaz.
	/* var x int16 = 128
	var y int8

	y = int8(x)

	fmt.Println(y) */

	/* x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// string bir ifadeyi type conversion ile integer bir değere dönüştüremeyiz
	fmt.Println(x + y) */

	// ASCII sayılarının karşılığına denk gelen karakterler string değere dönüştürülebilir.
	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)

}
