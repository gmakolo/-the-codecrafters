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
		fmt.Println("6.HELP")
		fmt.Scan(&op)
		if op == "5" {
			fmt.Println("Bye")
			break
		}

		switch op {
		case "1":
			for {
				fmt.Println("Enter the first number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break

			}

			for {
				fmt.Println("Enter the second number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break

			}

			fmt.Println("Result:", a+b)
		case "2":
			for {
				fmt.Println("Enter the first number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break
			}

			for {
				fmt.Println("Enter the second number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break

			}

			fmt.Println("Result:", a-b)
		case "3":
			for {
				fmt.Println("Enter the first number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break
			}
			for {
				fmt.Println("Enter the second number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break
			}
			fmt.Println("Result:", a*b)
		case "4":
			for {
				fmt.Println("Enter the first number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break
			}
			for {
				fmt.Println("Enter the second number: ")
				_, error := fmt.Scan(&a)
				if error != nil {
					fmt.Println("You have an error: ")
					fmt.Println(error)
					continue
				}
				break
			}
			if b != 0 {
				fmt.Println("Result:", a/b)

			}
			if b == 0 {
				fmt.Println("cannot divide by 0 :")
			}
		case "6":
			fmt.Println("How to use my calculator")
			fmt.Println("1. add")
			fmt.Println("2.sub")
			fmt.Println("3.multiply")
			fmt.Println("4.divide")
			fmt.Println("5.exist")

		default:
			fmt.Println("Unknown operation")

		}
	}
}
