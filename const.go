package main

import "fmt"

func main() {
	const name = "Sauki"
	fmt.Println("My name is", name)

	// ? Constant declaration boleh tidak digunakan
	const (
		age = 20
		isMarried = false

		// tidak digunakan
		halo string = "Halo"
		anjay string = "Anjay"
	)

	fmt.Println(anjay)

	// var (
	// 	keren = true
	// 	anjay string
	// )

	// anjay = "anjay"
	// fmt.Println("Anjay", anjay)


	fmt.Println("My age is", age)
	fmt.Println("Is married?", isMarried)
	// fmt.Println("Keren?", keren)
}