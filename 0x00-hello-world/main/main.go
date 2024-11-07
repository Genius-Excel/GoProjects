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

	str_result, num_idx_result := fizzbuzz.Fizzbuuz(30)

	// fmt.Println(result)

	for i := 0; i < len(num_idx_result); i++ {
		fmt.Printf("%v: %v", num_idx_result[i], str_result[i])
	}
}
