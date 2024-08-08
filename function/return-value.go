package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("Enter first number :")
	fmt.Scan(&num1)
	fmt.Println("Enter second number :")
	fmt.Scan(&num2)
	var result int = addNumbers(num1, num2)
	fmt.Println("result is", result)


	var avg = calcAvg(1, 2, 2, 1)
	fmt.Println("avg is", avg)
}

func addNumbers(num1 int, num2 int) int {

	return num1 + num2

}

// TODO : variadic params

func calcAvg(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}


	var avg = float64(total) / float64(len(numbers))

	return avg
}
