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

}
