package main

import "fmt"

func main() {
	var x, y float32
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&y)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f + %f = %f", x, y, x+y)
	case "-":
		fmt.Printf("%f - %f = %f", x, y, x-y)
	case "*":
		fmt.Printf("%f * %f = %f", x, y, x*y)
	case "/":
		if y == 0 {
			fmt.Print("Division by zero error")
		} else {
			fmt.Printf("%f / %f = %f", x, y, x/y)
		}
	default:
		fmt.Println("Invalid Operator")
	}
}
