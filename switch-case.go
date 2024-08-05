package main

import "fmt"

func main() {
	month := 1
	switch month {
	case 1:
		fmt.Println("January")
		fallthrough
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	default: 
		fmt.Println("Unknown")
	}
	
}