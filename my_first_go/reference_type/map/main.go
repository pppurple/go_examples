package main

import (
	"fmt"
)

func main() {
	basic()

	reteral()

	nest()

	reference()

	loop()
}

func basic() {
	m := make(map[int]string)

	m[1] = "apple"
	m[2] = "banana"
	m[100] = "cherry"
	fmt.Println(m)

	m[100] = "ddd"
	fmt.Println(m)
}

func reteral() {
	m := map[int]string{1: "apple", 2: "banana", 100: "cherry"}
	fmt.Println(m)

	m2 := map[int]string{
		1:  "America",
		2:  "Brazil",
		99: "Canada",
	}
	fmt.Println(m2)
}

func nest() {
	m := map[int][]int{
		1:  []int{1, 2, 3},
		2:  []int{4, 5, 6},
		88: []int{100, 101},
	}
	fmt.Println(m)

	m2 := map[int][]int{
		1:  {1, 2, 3},
		2:  {4, 5, 6},
		99: {1000, 1001},
	}
	fmt.Println(m2)
}

func reference() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m[1])
	fmt.Println(m[9])

	s, ok1 := m[1]
	fmt.Println(s)
	fmt.Println(ok1)
	s2, ok2 := m[9]
	fmt.Println(s2)
	fmt.Println(ok2)

	if _, ok := m[1]; ok {
		fmt.Println(ok)
	}
}

func loop() {
	m := map[int]string{
		1: "America",
		2: "Brazil",
		3: "Canada",
	}
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}

}
