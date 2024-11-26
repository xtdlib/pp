package main

import "github.com/xtdlib/pp"

type NestedExample struct {
	Name    string
	Age     int `pp:"-"` // This field will be skipped
	Email   string
	Address struct {
		City   string
		VV     []string
		State  string
		Zip    string `pp:"-"` // This field will be skipped
		Coords struct {
			Lat float64
			Lng float64
		}
	}
}

func main() {
	ex := NestedExample{
		Name:  "John Doe",
		Age:   30, // This won't be printed due to `pp:"-"` tag
		Email: "john.doe@example.com",
		Address: struct {
			City   string
			VV     []string
			State  string
			Zip    string `pp:"-"`
			Coords struct {
				Lat float64
				Lng float64
			}
		}{
			City:  "New York",
			VV:    []string{"a", "b", "c"},
			State: "NY",
			Zip:   "10001", // This won't be printed due to `pp:"-"` tag
			Coords: struct {
				Lat float64
				Lng float64
			}{
				Lat: 40.7128,
				Lng: -74.0060,
			},
		},
	}

	somevar := 3
	pp.Print(ex)
	pp.Print(somevar)
	pp.Print(3)
}
