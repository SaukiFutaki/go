package main

import "fmt"

func main() {
	var names [3	]string

	names = [3]string{"Sauki", "Rahmat", "Hidayat"}
	fmt.Println(names)


	// langsung
	var values = [3]int {
		90,
		95,
		80,
	}
	fmt.Println(values)

	// mengubah data di posisi index
	values[0] = 100
	fmt.Println(len(values))
	fmt.Println(values)


	var number = [...]int {
		1,
	} 

	fmt.Println(number)
}