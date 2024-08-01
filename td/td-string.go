package main

import "fmt"

func main(){
 var name string

 fmt.Print("Enter your name :")
 // ? Scan := untuk menerima inputan
 fmt.Scan(&name)
 fmt.Println("your name is", name)
 // ? len := untuk menghitung panjang string dari 0
 fmt.Println(len(name))
// ? [0] := untuk mengambil karakter pertama dari string di format ke ascii
 fmt.Println(name[0])
 // ? misal sauki maka akan menghasilkan s
 // TODO : lalu s di format ke ascii yaitu 115

}