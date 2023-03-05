package main

import "fmt"

var (
	legends     map[int]string
	calibration map[float64]bool
	table       map[string][]string
)

func main() {
	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514

	for key, val := range hist {
		adjVal := int(float64(val) * 0.100)
		fmt.Printf("%s (%d):", key, val)
		for i := 0; i < adjVal; i++ {
			fmt.Print(".")
		}
		fmt.Println()
	}
}
