package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}
}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}

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

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println((allusers))

	// if err := db.Delete("user", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
