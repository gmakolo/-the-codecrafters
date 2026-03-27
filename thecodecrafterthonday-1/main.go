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
		fmt.Println("5. exit")
		fmt.Scan(&op)
		if op == "5" {
			fmt.Println("bye")
			break
		}

		switch op {
		case "1":
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)

			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)

			fmt.Println("Result:", a+b)
		case "2":
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)

			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)

			fmt.Println("Result:", a-b)
		case "3":
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)

			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)

			fmt.Println("Result:", a*b)
		case "4":
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)

			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)
			if b != 0 {
				fmt.Println("Result:", a/b)

			}
			if b == 0 {
				fmt.Println("cannot divide by 0 :")
			}

		default:
			fmt.Println("Unknown operation")

		}
	}
}
