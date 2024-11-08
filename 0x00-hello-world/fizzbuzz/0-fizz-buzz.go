package fizzbuzz

func Fizzbuuz(num int) ([]string, []int) {
	// This function does the FizzBuzz program in Go.

	// IF the number is divisible by 3, return "Fizz"
	// IF the number is divisible by 5, return "Buzz"
	// IF the number is divisible by both 3 and 5, return "FizzBuzz"

	var str_slice []string
	var num_slice []int

	// Loop through the integer arguament provided.
	for i := 0; i <= num; i++ {
		string_element := ""

		if num%3 == 0 && num%5 == 0 {
			string_element = "Fizzbuzz"
		}
		if num%5 == 0 {
			string_element = "Buzz"
		}
		if num%3 == 0 {
			string_element = "Fizz"
		}

		str_slice = append(str_slice, string_element)
		num_slice = append(num_slice, i)
	}

	return str_slice, num_slice
	// if num%3 == 0 && num%5 == 0 {
	// 	return "FizzBuzz"
	// } else if num%5 == 0 {
	// 	return "Buzz"
	// } else if num%3 == 0 {
	// 	return "Fizz"
	// }

}
