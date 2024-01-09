package main

import (
	"encoding/json"
	"fmt"
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
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "1234567", "Google", Address{"San Francisco", "California", "United States", "94105"}},
		{"Emma", "29", "8743210", "Amazon", Address{"Seattle", "Washington", "United States", "98109"}},
		{"Carlos", "35", "7625894", "Microsoft", Address{"Redmond", "Washington", "United States", "98052"}},
		{"Aisha", "31", "9456732", "Apple", Address{"Cupertino", "California", "United States", "95014"}},
		{"Liam", "28", "6182947", "Facebook", Address{"Menlo Park", "California", "United States", "94025"}},
		{"Yuna", "26", "5371826", "Netflix", Address{"Los Gatos", "California", "United States", "95032"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}
}
