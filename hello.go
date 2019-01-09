package main

import "fmt"

func main() {
	fmt.Println("Hello Crazy Fucking World")
	foo()
	for i := 0 ;i < 15; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Im in fooo")
}

//control flow
//(1) sequence
//(2) loop; iterative
//(3) conditional
