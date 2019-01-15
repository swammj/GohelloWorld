package main

import (
	"fmt"
	"runtime"
)

var x bool

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
}
