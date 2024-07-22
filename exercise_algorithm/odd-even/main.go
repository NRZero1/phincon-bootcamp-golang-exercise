package main

import "fmt"

func main() {
	var number int
	var odd int
	for {
		fmt.Print("Input a number: ")
		_, err := fmt.Scanf("%d\n", &number)

		if err != nil {
			fmt.Println("Invalid input please enter a number")
			continue
		}

		odd = number % 2
		if odd != 0 {
			fmt.Printf("Number %d is odd", number)
		} else {
			fmt.Printf("Number %d is even", number)
		}
	}
}