package main

import "fmt"

func main() {
	var numb1,numb2 int
	var fruit string


	fmt.Print("Masukan angka ke 1: ")
	fmt.Scan(&numb1)

	fmt.Print("Masukan angka ke 2: ")
	fmt.Scan(&numb2)
	if numb1 > numb2 {
		fmt.Println("angka ke 1 :", numb1, "lebih besar dari angka ke 2 :", numb2)
	} else if numb2 > numb1 {
		fmt.Println("angka ke 2 :", numb2, "lebih besar dari angka ke 1 :", numb1)
	} else {
		fmt.Println("angka 1:", numb1, "sama dengan angka 2:", numb2)
	}


	fmt.Print("Masukan nama buah: ")
	fmt.Scan(&fruit)
	if fruit != "apel" {
		fmt.Println(fruit,", buah yang dimasukan bukan apel")
	} else {
		fmt.Println("buah yang dimasukan adalah apel")
	}



}