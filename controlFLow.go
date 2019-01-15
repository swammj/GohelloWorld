package main
//gobyexample.com
import "fmt"

func main() {

	//for init; condition; post
	for index := 0; index < 5; index++ {
		fmt.Println("Index is: ", index)
	}

	//nesting a loop
	for i := 0; i < 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("\t\tThe inner loop : %d\n", j)
		}
	}

	//for solo con condicion
	fmt.Println("***** FOR only with Condition")
	x := 1
	for (x<10)  {
		fmt.Println("X is: ",x)
		x++
	}


}
