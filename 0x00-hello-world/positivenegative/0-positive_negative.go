package positivenegative

import (
	"fmt"
	"math/rand"
)

func Postive_Negative() {

	n := rand.Intn(200)

	// Check if n is positive or negative
	if n < 0 {
		fmt.Printf("%d is a negative integer", n)
	} else if n == 0 {
		fmt.Printf("%d is zero", n)
	} else if n > 0 {
		fmt.Printf("%d is a positive integer", n)

	}

}
