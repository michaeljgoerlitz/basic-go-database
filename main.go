package main

import (
	"encoding/json"
)

type Address struct {
	City    string
	State   string
	Country string
	ZIP     json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {

}
