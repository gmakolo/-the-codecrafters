package main

import "fmt"

func main() {
	for {
		var op string
		var a, b float64

		fmt.Println("Enter the operation you want: ")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Scan(&op)

		fmt.Println("Enter the first number: ")
		fmt.Scan(&a)

		fmt.Println("Enter the second number: ")
		fmt.Scan(&b)

		switch op {
		case "1":
			fmt.Println("Result:", a+b)
		case "2":
			fmt.Println("Result:", a-b)
		case "3":
			fmt.Println("Result:", a*b)
		case "4":
			fmt.Println("Result:", a/b)
		default:
			fmt.Println("Unknown operation")
		}
	}
}
