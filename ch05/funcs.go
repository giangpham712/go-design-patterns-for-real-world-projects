package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf(
		"94 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(94),
	)

	for i := 0.0; i < 360.0; i += 45.0 {
		rad := func() float64 {
			return i * math.Pi / 180
		}()
		fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad)
	}
}
