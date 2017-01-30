package main

import "fmt"

func main() {
	// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
	// i := 1
	// for i <= 100 {
	// 	if i%3 == 0 {
	//
	// 		fmt.Println(i)
	// 	}
	// 	i++
	// }

	j := 1
	for j <= 100 {
		var fizzBuzz string
		if j%3 == 0 {
			fizzBuzz += "Fizz"
		}
		if j%5 == 0 {
			fizzBuzz += "Buzz"
		}

		if fizzBuzz != "" {
			fmt.Println(fizzBuzz)
		} else {
			fmt.Println(j)
		}
		j++
	}
}
