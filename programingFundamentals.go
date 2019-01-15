package main

import (
	"fmt"
	"runtime"
)

//var x bool
const (
	_  = iota// seinicializa en 1
	kb = 1 << (iota * 10) // lo incrementamos a lo que nosotros querramos
	mb 						//se incrementa por lo que esta arriba por eso no hay que ponerlo  aun que si se puede.
	gb
)

func main() {
	//*********BOOL********
	a := 7
	b := 29

	fmt.Println(a != b)

	//*********NUMERIC TYPES********
	//ALiAS byte = uint8
	// ALIAS rune = int32
	// just use float64

	//*********RUN TYPES********
	fmt.Println(runtime.GOOS)   //ver el tipo de OS en el que se trabaja
	fmt.Println(runtime.GOARCH) // ver el tipo de arquitectura de la compu

	//*********STRING TYPES********
	s := "This is a String"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	/*bs := []byte(s)
	fmt.Println(bs) //imprime el codigo ASCII de la letra
	fmt.Printf("%T", bs)*/

	//*********ITERATION********

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) //%#U es para ver el UTF 8 Code Point
	}

	fmt.Println("")
	for i, v := range s {
		fmt.Printf("Index position is %d we have %s \n", i, string(v)) // imprime letra por letra del string
	}

	letra := "H"
	fmt.Println(letra)

	bs := []byte(letra)
	fmt.Println(bs)

	//*********NUMERAL SYSTEM********
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)

	//*********BIT SHIFTING********
	fmt.Println()

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Println("*********IOTA to Increment********")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
