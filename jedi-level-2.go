package main

import (
	"fmt"
)

const (
	constante int = 27
	sintipo       = "Hola loco"
)

const(
	año = iota
	año2010 = (iota + 2009)
	año2011 //(iota + 2009)
	año2012 //(iota + 2009)
)

func main() {
	fmt.Printf("%s\n\n", "*******Excersise #1*******")
	numero := 27
	fmt.Printf("%d\n", numero)
	fmt.Printf("%b\n", numero)
	fmt.Printf("%#x\n", numero)

	fmt.Printf("\n%s\n\n", "*******Excersise #2*******")

	a := (27 == 27)
	b := (27 <= 24)
	c := (28 >= 27)
	d := (27 != 27)
	e := (27 < 28)
	f := (28 > 27)

	fmt.Println(a, b, c, d, e, f)

	fmt.Printf("\n%s\n\n", "*******Excersise #3*******")

	fmt.Println(constante, sintipo)

	fmt.Printf("\n%s\n\n", "*******Excersise #4*******")
	entero := 27
	fmt.Printf("%d\n", entero)
	fmt.Printf("%b\n", entero)
	fmt.Printf("%#x\n", entero)
	movio := entero << 1

	fmt.Println(movio)
	fmt.Printf("%b\n", movio)

	fmt.Printf("\n%s\n\n", "*******Excersise #5*******")

	cadema := `"Aqui esta esta cadena Jodida
				con todo y sus espacios
		Esta fresa"`

	fmt.Println(cadema)

	fmt.Printf("\n%s\n\n", "*******Excersise #6 IOTA*******")

	println(año2010,año2011 ,año2012 )









}
