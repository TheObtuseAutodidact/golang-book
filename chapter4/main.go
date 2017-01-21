package main

import "fmt"

func main() {
	// fmt.Print("Enter a number to magically see it doubled!: ")
	// var input float64
	// fmt.Scanf("%f", &input)
	//
	// output := input * 2
	//
	// fmt.Println(output)
	// fmt.Println("Tada!")

	// fmt.Print("Enter temp in Fahrenheit: ")
	// var fTemp float64
	// fmt.Scanf("%f", &fTemp)
	// fmt.Printf("%f in Celsius is %f", fTemp, fToC(fTemp))
	fmt.Print("Enter number of feet: ")
	var feet float32
	fmt.Scanf("%f", &feet)
	fmt.Printf("%f feet is equivalent to %f meters", feet, feetToMeters(feet))
}

func fToC(f float32) float32 {
	return (f - 32.0) * 5.0 / 9
}

func feetToMeters(f float32) float32 {
	return f * 0.3048
}
