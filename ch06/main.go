package main

import (
	"flag"
	"fmt"
	_ "fmt"
	_ "github.com/giangpham712/go-design-patterns-for-real-world-projects/ch06/volt"
)

var (
	op string
	v  float64
	r  float64
	i  float64
	p  float64
)

func init() {
	flag.Float64Var(&v, "v", 0.0, "Voltage value (volt)")
}

func main() {
	flag.Parse()

	fmt.Println(v)
}
