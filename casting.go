package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var number float64 = 10.9

	var bigNumber int64 = int64(number) // casting int32 to int64
	var floatNumber float32 = float32(bigNumber) // casting int64 to float32	

	fmt.Println(number)
	fmt.Println(bigNumber)
	fmt.Println(floatNumber)


	var isMarried bool = true
	// reflect.TypeOf() digunakan untuk mengetahui tipe data dari variabel
	fmt.Println(reflect.TypeOf(isMarried))

	//strconv.itoa > int to string

	var str  = strconv.Itoa(10)
	fmt.Println(reflect.TypeOf(str))


	//strconv.Atoi > string to int
	// fungsi atoi mengebalikan 2 nilai, nilai pertama adalah hasil konversi dan nilai kedua adalah error
	var num, _ = strconv.Atoi("10")
	fmt.Println(reflect.TypeOf(num))

}