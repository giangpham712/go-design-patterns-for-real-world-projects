package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perim() float64
}

type rect struct {
	name           string
	length, height float64
}

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perim() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
	return fmt.Sprintf(
		"%s[sides: a=%.2f b=%.2f c=%.2f]",
		t.name, t.a, t.b, t.c,
	)
}

func shapeInfo(s shape) string {
	return fmt.Sprintf(
		"Area = %.2f, Perim = %.2f",
		s.area(), s.perim(),
	)
}

func (r *rect) area() float64 {
	return r.length * r.height
}

func (r *rect) perim() float64 {
	return 2*r.length + 2*r.height
}

type polygon interface {
	shape
}

type curved interface {
	shape
	circonf()
}

type food interface {
	eat()
}

type veggie string

func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string

func (m meat) eat() {
	fmt.Println("Eating tasty", m)
}

func eat(f food) {
	veg, ok := f.(veggie)
	if ok {
		if veg == "okra" {
			fmt.Println("Yuk! not eating ", veg)
		} else {
			veg.eat()
		}
		return
	}

	mt, ok := f.(meat)
	if ok {
		if mt == "beef" {
			fmt.Println("Yuk! not eating ", mt)
		} else {
			mt.eat()
		}
		return
	}

	fmt.Println("Not eating whatever that is: ", f)
}

func eatWithSwitch(f food) {
	switch morse1 := f.(type) {
	case veggie:
		if morse1 == "okra" {
			fmt.Println("Yuk! not eating ", morse1)
		} else {
			morse1.eat()
		}
	case meat:
		if morse1 == "beef" {
			fmt.Println("Yuk! not eating ", morse1)
		} else {
			morse1.eat()
		}
	default:
		fmt.Println("Not eating whatever that is: ", f)
	}
}

func main() {
	r := &rect{"Square", 4.0, 4.0}
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{"Right Triangle", 1, 2, 3}
	fmt.Println(t, "=>", shapeInfo(t))

	var anyType interface{}
	anyType = 77.0
	anyType = "I am a string now"
	fmt.Println(anyType)

	printAnyType("The car is slow")
}

func printAnyType(val interface{}) {
	fmt.Println(val)
}
