package main

import (
	"encoding/json"
	"fmt"
)

type diameter int

type name struct {
	long   string
	short  string
	symbol rune
}

type planet struct {
	diameter
	name
	desc string
}

type Person struct {
	Name    string `json:"person_name"`
	Title   string `json:"person_title"`
	Address `json:"person_address_obj"`
}

type Address struct {
	Street string `json:"person_addr_street"`
	City   string `json:"person_city"`
	State  string `json:"person_state"`
	Postal string `json:"person_postal_code"`
}

var (
	empty    struct{}
	car      struct{ make, model string }
	currency struct {
		name, country string
		code          int
	}
	node struct {
		edges  []string
		weight int
	}
	person struct {
		name    string
		address struct {
			street      string
			city, state string
			postal      string
		}
	}
)

func main() {
	currency = struct {
		name, country string
		code          int
	}{
		"USD", "United States",
		840,
	}
	fmt.Println(currency)

	car = struct{ make, model string }{make: "Ford", model: "F150"}

	earth := planet{
		diameter: 7926,
		name: name{
			long:   "Earth",
			short:  "E",
			symbol: '\u2641',
		},
		desc: "Third rock from the Sun",
	}

	fmt.Println(earth)

	saturn := planet{}
	saturn.diameter = 120536
	saturn.long = "Saturn"
	saturn.short = "S"
	saturn.symbol = '\u2644'
	saturn.desc = "Slow mover"

	p := Person{
		Name:  "Vladimir Vivien",
		Title: "Author",
	}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
