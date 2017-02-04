package main

// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

import "fmt"

func main() {
	a, b := 1, 3
	fmt.Println(swap(&a, &b))

	c, d := 11111, 99999
	fmt.Println(swap(&c, &d))
}
func swap(x, y *int) (int, int) {
	x, y = y, x
	return *x, *y
}
