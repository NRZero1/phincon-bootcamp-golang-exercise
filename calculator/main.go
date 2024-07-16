package main

import (
	"calculator/model"
	"fmt"
)

func main() {
	var menu int;
	var value1 float64
	var value2 float64
	var loopAgain string
	
	for ok := true ; ok; ok = (loopAgain == "y" || loopAgain == "Y") {
		fmt.Println()
		fmt.Println("Calculator Program")
		fmt.Println("=========================")
		fmt.Println("Menu: ")
		fmt.Println("1. Addition")
		fmt.Println("2. Substraction")
		fmt.Println("3. Division")
		fmt.Println("4. Multiplication")
		fmt.Println()
		fmt.Print("Input menu: ")
		receivedInput, err := fmt.Scanf("%d\n", &menu)
		fmt.Println()

		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		if receivedInput == 5 {
			fmt.Println("Invalid number inputted, please enter a number between 1-4")
			continue
		}
		
		fmt.Print("Input value 1: ")
		_, errValue1 := fmt.Scanf("%f\n", &value1)


		fmt.Print("Input value 2: ")
		_, errValue2 := fmt.Scanf("%f\n", &value2)

		if errValue1 != nil {
			fmt.Println("Invalid value 1 inputted, please enter a number")
			continue
		}

		if errValue2 != nil {
			fmt.Println("Invalid value 2 inputted, please enter a number")
			continue
		}

		calculator := model.Calculator{}
		calculator.SetVal1(value1)
		calculator.SetVal2(value2)

		defer func() {
			fmt.Println("Recovered from panic: ", recover())
		}()

		intValue, floatValue := calculator.Calculate(menu)

		fmt.Println("Result in integer: ", intValue)
		fmt.Println("Result in float: ", floatValue)

		fmt.Print("Run program again? (y/n) ")
		fmt.Scanf("%s\n", &loopAgain)
	}
}