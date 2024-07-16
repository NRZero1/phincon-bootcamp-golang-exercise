package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	goods map[string]string
}

func getAllProduct(listOfProduct []Product) {
	for _, product := range listOfProduct {
		for k, v := range product.goods {
			fmt.Println("Name: ", k)
			fmt.Println("Stock: ", v)
		}
	}
}

func handleError() {
	fmt.Println("Recover: ", recover())
}

func main() {
	var listOfProduct []Product
	defer handleError()

	for {
		fmt.Println()
		fmt.Println("Program Simple Stock Management System")
		fmt.Println("Stock saat ini: ")

		if (len(listOfProduct) == 0) {
			fmt.Println("Product is empty")
		} else {
			getAllProduct(listOfProduct)
		}

		var name string
		var inputtedStock int
		
		fmt.Println()
		fmt.Println("Input Product: ")
		fmt.Print("Product name: ")
		_, errName := fmt.Scanf("%s\n", &name)
		fmt.Print("Product stock: ")
		_, errStock := fmt.Scanf("%d\n", &inputtedStock)

		if (errName != nil) {
			panic("An error occured with inputted name")
		}

		if (errStock != nil) {
			panic("An error occured with inputted stock")
		}

		product := Product{
			goods: map[string]string {
				name: strconv.Itoa(inputtedStock),
			},
		}
		listOfProduct = append(listOfProduct, product)
	}
}