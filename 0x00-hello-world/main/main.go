package main

import (
	"fmt"

	"github.com/Genius-Excel/GoProjects/0x00-hello-world/fizzbuzz"
	"github.com/Genius-Excel/GoProjects/0x00-hello-world/positivenegative"
)

func main() {

	// This is the main function to call the helper function

	positivenegative.Postive_Negative()
	fmt.Println("Fizz Buzz Question")

	result := fizzbuzz.Fizzbuuz(30)

	fmt.Println(result)

	// Tesiting how Arrays work
	my_array := [...]int{5, 8, 9, 0}

	for i := 0; i < len(my_array); i++ {
		fmt.Println(my_array[i])
	}

	// fmt.Printf("Antoehr type of loop with index")
	// for value := range len(my_array)  {
	// 	fmt.Printf("%v\n", my_array[value])
	// }
}
