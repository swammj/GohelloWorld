package main

import "fmt"

//outside functions declarations
var y = 50

var nombre int //Declaracion de variable sin asignacion

var cadena string = "Esta es una cadena"

func main() {
	//DECLARE a variable and ASSIGNING A VALUE with :=
	x := 34
	fmt.Println(x)
	//ASSIGNING a value to de DECLARED variable
	x = 93
	fmt.Println(x)
	xq := "Hola"
	fmt.Println(xq, x)
	nombre = 40
	fmt.Println("Variable nombre es :", nombre, "And y is: ", y)

	//PRINT F
	entero := 40
	fmt.Printf("%T,\n", entero)

	fmt.Printf("%s\n", cadena)

	cadena2 := `Hola "dijo ella"` //fresa!!!
	fmt.Println(cadena2)

	//TIPOS
	type chocolate int
	var myTipo chocolate = 23498

	fmt.Println(myTipo)
	fmt.Printf("%T\n", myTipo)

	//convertion

	var convertido int = 40

	convertido = int(myTipo)
	fmt.Println(convertido)
	fmt.Printf("%T\n", myTipo)

}
