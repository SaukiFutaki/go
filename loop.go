package main

import "fmt"

func main() {

	
	for num := 0; num <= 50; num++ {
		 if num % 2 == 0 {
			fmt.Println(num)	
		 }
	}


	for _,value := range []int{1,2,3,4,5} {
		fmt.Println(value)
	}

}