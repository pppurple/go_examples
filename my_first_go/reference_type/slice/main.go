package main

import "fmt"

func main() {
	// slice
	s := make([]int, 10)
	fmt.Println(s)

	// array
	var a [10]int
	fmt.Println(a)

	// len
	fmt.Println(len(s))

	// cap
	fmt.Println("***cap")
	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// reteral
	fmt.Println("***reteral")
	sr := []int{1, 2, 3, 4, 5}
	fmt.Println(sr)

	// simple slice expressions
	as := [5]int{1, 2, 3, 4, 5}
	ss := as[0:2]
	fmt.Println(ss)

	// string
	str := "ABCDEF"[1:3]
	fmt.Println(str)
}
