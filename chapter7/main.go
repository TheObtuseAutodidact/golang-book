package main

import "fmt"

func main() {
	fmt.Println(half(3))
	fmt.Println(half(4))
	fmt.Println(greatest(1, 2, 3, 4))
	fmt.Println(greatest(4, 3, 2, 1))
	fmt.Println(greatest(2, 17, 10987, 8, 63))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
}

// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
func half(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
func greatest(i ...int) int {
	biggest := i[0]
	for j := range i {
		if i[j] > biggest {
			biggest = i[j]
		}
	}
	return biggest
}

// Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
func makeOddGenerator() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
