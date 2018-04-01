package main

import (
	"fmt"
)

func main() {
	address()

	dereference()

	incMain()

	powMain()

	arrayPointer()

	lenCap()

	loop()
}

func address() {
	var i int
	p := &i
	fmt.Printf("%T\n", p)
	pp := &p
	fmt.Printf("%T\n", pp)
}

func dereference() {
	var i int
	p := &i
	i = 5
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)
}

func incMain() {
	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)
}

func inc(p *int) {
	*p++
}

func powMain() {
	p := &[3]int{1, 2, 3}
	pow(p)
	fmt.Println(p)
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func arrayPointer() {
	a := [3]string{"america", "brazil", "canada"}
	p := &a
	fmt.Println(a[1])
	fmt.Println(p[1])
	p[2] = "Dan"
	fmt.Println(a[2])
	fmt.Println(p[2])
}

func lenCap() {
	p := &[3]int{1, 2, 3}
	fmt.Println(len(p))
	fmt.Println(cap(p))
	fmt.Println(p[0:2])
}

func loop() {
	p := &[3]string{"America", "Brazil", "Canada"}

	for i, v := range p {
		fmt.Println(i, v)
	}
}
