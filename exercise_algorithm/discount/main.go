package main

import "fmt"

func handleError(message string) {
	fmt.Println("Error message: ", message)
}

func main() {
	var price = 0
	var total = 0
	var qty = 0
	var discountInput string

	for {
		fmt.Print("Please enter a price: ")
		_, errPrice := fmt.Scanf("%d\n", &price)

		if errPrice != nil {
			handleError("Invalid price input, please enter a number")
			continue
		}

		fmt.Print("\nPlease enter quantity: ")
		_, errQty := fmt.Scanf("%d\n", &qty)

		if (errQty != nil) {
			handleError("Invalid qty input, please enter a number")
			continue
		}

		fmt.Print("Is there a discount? (y/n) ")
		_, errDiscount := fmt.Scanf("%s\n", &discountInput)

		if errDiscount != nil || (discountInput != "y" && discountInput != "n") {
			handleError("Invalid discount input, please enter either y or n")
			continue
		}

		discount := discountInput == "y"

		if discount {
			total = (qty * price) - (((qty * price) * 10) / 100)
		} else {
			total = qty * price
		}

		fmt.Printf("Total is %d \n", total)
	}
}