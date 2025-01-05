package main

import (
	"errors"
	"fmt"
)

func addition(x, y float64) float64 {
	return x + y
}

func subtraction(x, y float64) float64 {
	return x - y
}

func multiplication(x, y float64) float64 {
	return x * y
}

func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return x / y, nil
}

func help() {
	fmt.Println("Welcome \n 1. Addition \n 2. Subtraction \n 3. Multiplication \n 4. Division")
}

func main() {
	fmt.Println("WELCOME TO CALCULATOR APP.")

	fmt.Println("\n 1. Addition \n 2. Subtraction \n 3. Multiplication \n 4. Division \n 5. Help ")

	var operation int
	fmt.Print("Enter operation number: ")
	fmt.Scan(&operation)

	// switch operation
	switch operation {
	case 1:
		var number1 float64
		var number2 float64

		// allow user to enter first number and second
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)

		fmt.Println(addition(number1, number2))
	case 2:
		var number1 float64
		var number2 float64

		// allow user to enter first number and second
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)

		fmt.Println(subtraction(number1, number2))
	case 3:
		var number1 float64
		var number2 float64

		// allow user to enter first number and second
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)

		fmt.Println(multiplication(number1, number2))
	case 4:
		var number1 float64
		var number2 float64

		// allow user to enter first number and second
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)

		result, err := division(number1, number2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	case 5:
		help()
	default:
		fmt.Println("Unknown operation!")
	}
}
