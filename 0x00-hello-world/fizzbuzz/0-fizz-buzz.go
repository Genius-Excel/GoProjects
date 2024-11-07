package fizzbuzz

func Fizzbuuz(num int) string {
	// This function does the FizzBuzz program in Go.

	// IF the number is divisible by 3, return "Fizz"
	// IF the number is divisible by 5, return "Buzz"
	// IF the number is divisible by both 3 and 5, return "FizzBuzz"

	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	} else if num%5 == 0 {
		return "Buzz"
	} else if num%3 == 0 {
		return "Fizz"
	}

	return "Not an integer"
}
