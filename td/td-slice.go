package main

import "fmt"

func main() {
	// point of slice := pointer, length, capacity
	// length !> capacity

	months := [...]string{
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
	// fmt.Println(months)
	slice := months[4:7] // -> index 4,5,6 //TODO : pointer : 4, length : 3, capacity : 8 karena dari index 4 sampai 11
	// TODO : pointer : penunjuk data pertama
	// TODO : length : panjang data jadi klo pointer 4 dan length 3 maka data yang diambil 4,5,6
	fmt.Println(slice[2]) // Juli
	// fmt.Println(slice[3]) // akan error karena index 3 tidak ada di slice

	slice2 := months[6:] // -> index 6 sampai 11
	slice3 := months[:6] // -> index 0 sampai 5
	slice4 := months[:]  // -> index 0 sampai 11
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(slice4)

	// ? jika dibuat secara manual
	// var slice5 []string = months[4:7] // -> bedanya apa dengan array, klo di slice kita tidak perlu menuliskan panjang array

	// ! function slice
	fmt.Println(len((slice2)))          // ->
	fmt.Println(cap((slice)))           // -> 8 karena dari index 4 sampai 11
	fmt.Println(append(slice, "Sauki")) // -> menambahkan data di slice

	// TODO : APPEND
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
	}
	fmt.Println(days)
	daySlice := days[1:3]
	fmt.Println(daySlice)
	daySlice[0] = "Senin"
	daySlice[1] = "Selasa"
	fmt.Println(daySlice)
	fmt.Println(days) // -> array days juga berubah karena slice merupakan reference dari array days
	daySlice2 := append(daySlice, "Jumat")
	fmt.Println(daySlice2) // -> menambahkan data di slice, // ? jika data di slice diubah maka data di array juga berubah
	fmt.Println(days)      // -> array days juga berubah karena slice merupakan reference dari array days
	fmt.Println(append(daySlice, "Sauki"))
	fmt.Println(days) // -> array days juga berubah karena slice merupakan reference dari array days

	// TODO : MAKE
	newSlice := make([]string, 2, 5) // -> make([]typeData, length, capacity)
	newSlice[0] = "Sauki"
	newSlice[1] = "Rahmat"
	// newSlice[2] = "Budi" // -> akan error karena panjang slice hanya 2
	newSlice2 := append(newSlice, "Budi")
	fmt.Println(len(newSlice2))
}
