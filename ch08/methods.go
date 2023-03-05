package main

import "fmt"

type gallon float64
type ounce float64
type cup float64
type quart float64

func (g gallon) quart() quart {
	return quart(g * 4)
}

func (o ounce) cup() cup {
	return cup(o * 0.1250)
}

func (c cup) quart() quart {
	return quart(c * 0.25)
}

func (c cup) ounce() ounce {
	return ounce(c * 8.0)
}

func (q quart) cup() cup {
	return cup(q * 4.0)
}

func (g gallon) half() {
	g = gallon(g * 0.5)
}

func (g *gallon) double() {
	*g = gallon(*g * 2)
}

func main() {
	gal := gallon(5)
	fmt.Printf("%.2f gallons = %.2f quarts\n", gal, gal.quart())
	ozs := gal.quart().cup().ounce()
	fmt.Printf("%.2f gallons = %.2f ounces\n", gal, ozs)

	gal = 5
	gal.half()
	fmt.Println(gal)
	gal.double()
	fmt.Println(gal)
}
