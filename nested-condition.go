package main

import "fmt"

func main() {
	var (
		score int
		name  string
	)

	fmt.Print("Masukan nama : ")
	fmt.Scan(&name)

	fmt.Print("Masukan nilai : ")
	fmt.Scan(&score)

	if name == "sauki" {
		if score >= 80 {
			fmt.Println("Selamat", name, "anda lulus")
		}
	} else {
		fmt.Println("Maaf", name, "anda tidak lulus karena kami hanya meluluskan orang yang bernama sauki")
	}
}
