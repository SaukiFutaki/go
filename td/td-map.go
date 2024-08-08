package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Sauki"
	// person["title"] = "Programmer"
	// person["address"] = "Jl. Jendral Sudirman"
	// fmt.Println(person)
	var person = map[string]string{ // -> [string] adalah key, string adalah value
		"name": "Sauki",
		"title": "Programmer",
		"address": "Jl. Jendral Sudirman",
	
	
	}
	fmt.Println(person)
}