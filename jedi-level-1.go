package main

import "fmt"

var a int = 43
var b string = "Swamy Jose"
var c bool = true

type myTipo int

var hola myTipo

var yu int

func main() {
	fmt.Println("************Exercise 1********")
	x := 42
	y := "James Bond"
	z := true


	fmt.Println("************Exercise 2********")
	fmt.Println(x, "\t", "\t", y, "\t", z)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("************Exercise 3********")
	fmt.Println(a, b, c)
	cadema := fmt.Sprintf("%v%v%v", a, b, c)
	println(cadema)

	fmt.Println("************Exercise 4********")
	fmt.Println(hola)
	fmt.Printf("%T\n", hola)
	hola = 42
	fmt.Println(hola)

	fmt.Println("************Exercise 5********")
	yu = int(hola)
	fmt.Println(y)
	fmt.Printf("%T\n", yu)

	fmt.Println("************Exercise 6********")



}
