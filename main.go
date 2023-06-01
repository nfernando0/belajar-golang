package main

import (
	"fmt"
)

func main() {
	
	var num1, num2 float64
	var operator string

	fmt.Println("Enter a number : ")
	fmt.Scanln(&num1)

	fmt.Println("Enter a number : ")
	fmt.Scanln(&num2)

	fmt.Println("Enter a operator : ")
	fmt.Scanln(&operator);

	switch operator {
	case "+" :
		fmt.Printf("%.3f %s %.3f = %.3f", num1, operator, num2, num1 + num2)
	case "-" :
		fmt.Printf("%.3f %s %.3f = %.3f", num1, operator, num2, num1 - num2)
	case "*" :
		fmt.Printf("%.3f %s %.3f = %.3f", num1, operator, num2, num1 * num2)
	case "/" :
		if(num2 == 0) {
			fmt.Println("Divide by zero situation")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f", num1, operator, num2, num1 / num2)
		}

	default:
		fmt.Println("Invalid Operator")

	}

}