package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		var command string
		var number string
		var base string

		fmt.Print("What would you like to convert?")
		fmt.Scan(&command)

		if command == "quit" {
			break
		}

		if command != "convert" {
			fmt.Println("Unknown command")
			continue
		}

		fmt.Scan(&number, &base)

		if base == "hex" {

			decimal, err := strconv.ParseInt(number, 16, 64)

			if err != nil {
				fmt.Println("Invalid hex")
				continue
			}

			fmt.Println("✦ Decimal:", decimal)
		}

		if base == "bin" {

			decimal, err := strconv.ParseInt(number, 2, 64)

			if err != nil {
				fmt.Println("Invalid binary")
				continue
			}

			fmt.Println("✦ Decimal:", decimal)
		}

		if base == "dec" {

			decimal, err := strconv.ParseInt(number, 10, 64)

			if err != nil {
				fmt.Println("Invalid decimal")
				continue
			}

			binary := strconv.FormatInt(decimal, 2)
			hex := strings.ToUpper(strconv.FormatInt(decimal, 16))

			fmt.Println("✦ Binary:", binary)
			fmt.Println("✦ Hex:", hex)
		}

		if base != "hex" && base != "bin" && base != "dec" {
			fmt.Println("Base must be hex, bin, or dec")
		}
	}
}
