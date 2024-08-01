package main

import "fmt"

func main(){
 var name string

 fmt.Print("Enter your name :")
 // ? Scan := untuk menerima inputan
 fmt.Scan(&name)
 fmt.Println("your name is", name)

}