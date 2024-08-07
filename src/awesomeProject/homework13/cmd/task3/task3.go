package main

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

type example struct {
	a []int
	b bool
	c int32
	d string
}

func main() {
	ex := example{
		a: []int{1, 2, 3}, // слайс - это структура, занимающая 24 байта + 24 байта за 3 значения = 48
		b: true,           // 1 байт
		d: "1234",         // 16 байт строка + 4 байта 4 символа = 20 байт
	} // 48 + 1 + 20 = 69, также не забудем, что структура example имеет с int32 (4 байта), если не задавать значение,
	// значит будет 0 по умолчанию => память выделится => 69 + 4 байта = 73 байта
	fmt.Println("Размер в байтах для ex:", size.Of(ex))
	// получится 76 байт из-за struct padding под машинное слово длиной 4 байта 73 + 3 = 76

	ex1 := example{
		a: []int{1, 2, 3}, // 48
		b: true,           // 1
		d: "1234",         // 20
		c: 100,            // 4
	} // тут тоже самое
	fmt.Println("Размер в байтах для ex1:", size.Of(ex1))
	// получится 76 байт из-за struct padding под машинное слово длиной 4 байта 73 + 3 = 76
}
