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

	// append
	sa := []int{1, 2, 3}
	sa = append(sa, 4)
	fmt.Println(sa)
	sa = append(sa, 5, 6, 7)
	fmt.Println(sa)

	sa1 := []int{1, 2, 3}
	sa2 := []int{4, 5, 6}
	sa3 := append(sa1, sa2...)
	fmt.Println(sa3)

	// auto expand slice size
	autoExpand()

	// copy
	copySlice()

	// full slice expressions
	fullSliceExpression()

	// loop
	sliceLoop()

	// variable args
	variableArgs()

	// reference type
	arr := [3]int{1, 2, 3}
	double(arr) // call by value
	fmt.Println(arr)
	sli := []int{1, 2, 3}
	doubleRef(sli) // call by reference
	fmt.Println(sli)

}

func autoExpand() {
	s := make([]int, 0, 0)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 5)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 6, 7, 8, 9)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func copySlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	n := copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(n)
}

func fullSliceExpression() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := a[2:4]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := a[2:4:4]
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := a[2:4:6]
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
}

func sliceLoop() {
	s := []string{"apple", "banana", "cherry"}

	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}

func variableArgs() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())

	s := []int{2, 4, 6}
	fmt.Println(sum(s...))
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// call by value
func double(a [3]int) {
	for i, v := range a {
		a[i] = 2 * v
	}
	return
}

// call by reference
func doubleRef(a []int) {
	for i, v := range a {
		a[i] = 2 * v
	}
	return
}
