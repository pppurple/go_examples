package main

import (
	"fmt"
)

func main() {
	basicError()

	toStringInterface()

	stringer()
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "error occured.", ErrCode: 500}
}

func basicError() {
	// type error
	err := RaiseError()
	fmt.Println(err.Error())

	// type MyError
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.Error())
	}
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("%s[%s]", c.Model, c.Number)
}

func toStringInterface() {
	vs := []Stringify{
		&Person{Name: "Alice", Age: 20},
		&Car{Number: "1111-2222", Model: "classic"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}

type T struct {
	Id   int
	Name string
}

type T2 struct {
	Id   int
	Name string
}

func (t2 *T2) String() string {
	return fmt.Sprintf("<<%d, %s>>", t2.Id, t2.Name)
}

func stringer() {
	t := &T{Id: 10, Name: "Apple"}
	fmt.Println(t)

	t2 := &T2{Id: 20, Name: "Banana"}
	fmt.Println(t2)
}
