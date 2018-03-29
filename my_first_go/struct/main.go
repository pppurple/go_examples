package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

func main() {
	typeMain()

	defineStruct()

	nest()
	nest2()
	nest3()

	pointer()

	newStruct()

	method()

	aliasMethod()

	constructor()

	methodAsFunc()

	pointerTypeReciever()

	sliceStruct()

	mapStruct()

	tag()

	tagJson()
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

	pt2 := Point{X: 1, Y: 2}
	fmt.Println(pt2.X)
	fmt.Println(pt2.Y)

	pt3 := Point{Y: 22}
	fmt.Println(pt3.X)
	fmt.Println(pt3.Y)
}

type Feed struct {
	Name string
	Age  uint
}

type Animal struct {
	Name string
	Feed Feed
}

func nest() {
	a := Animal{
		Name: "Dog",
		Feed: Feed{
			Name: "dog food",
			Age:  2,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Feed.Age)

	a.Feed.Age = 199
	fmt.Println(a.Feed.Age)
}

type Feed2 struct {
	Name string
	Age  uint
}

type Animal2 struct {
	Name string
	Feed2
}

func nest2() {
	a := Animal2{
		Name: "cat",
		Feed2: Feed2{
			Name: "tsuna",
			Age:  22,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed2.Name)
	fmt.Println(a.Age)

	a.Age = 100
	fmt.Println(a.Age)
}

type Base struct {
	Name string
	Age  int
}

type A struct {
	Base
	Country    string
	Population int
}

type B struct {
	Base
	Company  string
	Location string
}

func nest3() {
	a := A{
		Base:       Base{Name: "Alice", Age: 20},
		Country:    "America",
		Population: 10000,
	}
	fmt.Println(a.Name)
	fmt.Println(a.Age)
	fmt.Println(a.Country)
	fmt.Println(a.Population)

	b := B{
		Base:     Base{Name: "Bob", Age: 32},
		Company:  "IBM",
		Location: "America",
	}
	fmt.Println(b.Name)
	fmt.Println(b.Age)
	fmt.Println(b.Company)
	fmt.Println(b.Location)
}

func swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func pointer() {
	p := Point{X: 1, Y: 2}
	swap(&p)

	fmt.Println(p.X)
	fmt.Println(p.Y)
}

type Person struct {
	Name    string
	Age     int
	Country string
}

func newStruct() {
	p := new(Person)

	p.Name = "Bobby"
	p.Age = 29
	p.Country = "China"

	fmt.Println(*p)
}

type Double struct{ X, Y int }

func (d *Double) Debug() {
	fmt.Printf("<%d, %d>\n", d.X, d.Y)
}

func (d *Double) Sum(pd *Double) int {
	return d.X + pd.X + d.Y + pd.Y
}

func method() {
	d := &Double{X: 1, Y: 99}
	d.Debug()

	sum := d.Sum(&Double{X: 2, Y: 3})
	fmt.Println(sum)
}

type IntPoint struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *IntPoint) Distance(dp *IntPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}
func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(x*x + y*y)
}

func aliasMethod() {
	fmt.Println(MyInt(4).Plus(2))

	ip := IntPair{5, 9}
	fmt.Println(ip.First())
	fmt.Println(ip.Second())

	s := Strings{"A", "B", "C"}.Join(":")
	fmt.Println(s)
}

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Second() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	joined := ""
	for _, v := range s {
		if joined != "" {
			joined += d
		}
		joined += v
	}
	return joined
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	u := new(User)
	u.Name = name
	u.Age = age
	return u
}

func constructor() {
	fmt.Println(NewUser("Ritchie", 54))
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d,%d]\n", p.X, p.Y)
}

func methodAsFunc() {
	f := (*Point).ToString
	s := f(&Point{X: 10, Y: 20})
	fmt.Println(s)
}

type P struct{ X, Y int }

func (p P) Set(x, y int) {
	p.X = x
	p.Y = y
}

type PP struct{ X, Y int }

func (p *PP) Set(x, y int) {
	p.X = x
	p.Y = y
}

func pointerTypeReciever() {
	p1 := P{}
	p1.Set(1, 2)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p2 := &P{}
	p2.Set(3, 4)
	fmt.Println(p2.X)
	fmt.Println(p2.Y)

	pp := PP{}
	pp.Set(5, 6)
	fmt.Println(pp.X)
	fmt.Println(pp.Y)

	pp2 := &PP{}
	pp2.Set(7, 8)
	fmt.Println(pp2.X)
	fmt.Println(pp2.Y)
}

type PT struct{ X, Y int }
type PTs []*PT

func sliceStruct() {
	ps := make([]PT, 5)
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}

	pts := PTs{}
	pts = append(pts, &PT{X: 1, Y: 2})
	pts = append(pts, nil)
	pts = append(pts, &PT{X: 4, Y: 7})
	fmt.Println(pts.ToString())
}

func (pts PTs) ToString() string {
	str := ""
	for _, p := range pts {
		if str != "" {
			str += ":"
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}

func mapStruct() {
	m1 := map[User]string{
		{Name: "alice", Age: 10}: "America",
		{Name: "bobby", Age: 20}: "Canada",
	}
	m2 := map[int]User{
		1: {Name: "alice", Age: 10},
		2: {Name: "bobby", Age: 20},
	}
	fmt.Println(m1)
	fmt.Println(m2)
}

type Color struct {
	Id   int    "color id"
	Name string "color name"
}

func tag() {
	c := Color{Id: 1, Name: "orange"}

	t := reflect.TypeOf(c)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Type, f.Name, f.Tag)
	}
}

type Country struct {
	Id         int    `json:"c_id"`
	Name       string `json:"c_name"`
	Population int    `json:"c_pop"`
}

func tagJson() {
	c := Country{Id: 1, Name: "America", Population: 10000}
	bs, _ := json.Marshal(c)
	fmt.Println(string(bs))
}
