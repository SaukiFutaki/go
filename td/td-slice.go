package main

import "fmt"

func main() {
	// point of slice := pointer, length, capacity
	// length !> capacity

	month := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei", // -> index 4
		"Juni",
		"Juli",
		"Agustus", // -> index 7
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(month)
	slice:= month[4:7] // -> index 4,5,6 //TODO : pointer : 4, length : 3, capacity : 8 karena dari index 4 sampai 11
	// TODO : pointer : penunjuk data pertama
	// TODO : length : panjang data jadi klo pointer 4 dan length 3 maka data yang diambil 4,5,6
	fmt.Println(slice[2]) // Juli
	fmt.Println(slice[3]) // akan error karena index 3 tidak ada di slice
}
