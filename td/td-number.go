package main

import "fmt"

func main() {
	fmt.Println(2)

	// ? Augmented assignment
	// ? +=, -=, *=, /=, %=
	var num int

	fmt.Print("Masukan angka(akan ditambah 10) : ")
	fmt.Scan(&num)
	num += 10
	fmt.Println(num)
}
