package main

import (
	"fmt"
)

func main() {
	typeMain()

	defineStruct()
}

func typeMain() {
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)
	fmt.Println(n2)

	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float32
		IntsChannel chan []int
	)
	pair := IntPair{1, 3}
	strs := Strings{"AA", "BB", "CC"}
	amap := AreaMap{"America": {12.0, 29.0}}
	ich := make(IntsChannel)

	fmt.Println(pair)
	fmt.Println(strs)
	fmt.Println(amap)
	fmt.Println(ich)

	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n)
}

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

type Point struct {
	X int
	Y int
}

func defineStruct() {

	var pt Point
	fmt.Println(pt.X)
	fmt.Println(pt.Y)

	pt.X = 11
	pt.Y = 22
	fmt.Println(pt.X)
	fmt.Println(pt.Y)
}

func compositeLiterals() {
	pt := Point{1, 2}
	fmt.Println(pt.X)
	fmt.Println(pt.Y)

}
