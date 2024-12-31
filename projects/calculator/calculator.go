package main

import "fmt"

// addition
func addition(x, y float64)float64  {
	return x + y
}

// subtraction
func subtraction(x, y float64)float64  {
	return x / y
}

// multiplication
func multiplication(x, y float64)float64  {
	return x * y
}

// division
func division(x, y float64)float64  {
	if y != 0 {
		return x / y
	} else{
		return 0
	}
}
func main(){
	fmt.Println("WELCOME TO CALCULATOR APP.")

	fmt.Println("\n 1. Addition \n 2. Subtraction \n 3. Multiplication \n 4. Division")

	var operation string
	fmt.Print("Enter operation number: ")
	fmt.Scan(&operation)
}