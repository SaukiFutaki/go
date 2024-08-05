package main

import "fmt"

func main() {
	// int8 : -128 to 127
	// int16 : -32768 to 32767
	// int32 : -2147483648 to 2147483647
	// int64 : -9223372036854775808 to 9223372036854775807
	//uint8 : 0 to 255
	//uint16 : 0 to 65535
	//uint32 : 0 to 4294967295
	//uint64 : 0 to 18446744073709551615
	// float32 : 1.18E-38 to 3.4E+38
	// float64 : 2.23E-308 to 1.8E+308 

	var num32 int32 
	num32 = 2
	fmt.Println(num32)
	fmt.Println(2)

	// ? Augmented assignment
	// ? +=, -=, *=, /=, %=
	var num int

	fmt.Print("Masukan angka(akan ditambah 10) : ")
	fmt.Scan(&num)
	num += 10
	fmt.Println(num)
}
