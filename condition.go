package main

import "fmt"

func main() {
	var numb1,numb2,score1,score2 int
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


	fmt.Print("Masukan nilai 1: ")
	fmt.Scan(&score1)

	fmt.Print("Masukan nilai 2:")
	fmt.Scan(&score2)

	if  score1 >= 80 && score2 >= 80  {
		fmt.Println("Lulus")
	} else if score1 >= 80 || score2 >= 80 {
		fmt.Println("Remidi")
	} else {
		fmt.Println("Tidak Lulus")
	}

	

}