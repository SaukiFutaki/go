package main

import "fmt"

func main() {
	type name string
	var nama name = "Sauki"
	fmt.Println(nama)
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var i int = 10
	var j int = 1
	fmt.Println(i) // 10
	i++
	fmt.Println(i) // 11
	fmt.Println(j) // 1
	j += i	
	fmt.Println(j) // 12
}
