package main

import "fmt"

func main() {
	var score int
	for {
		fmt.Print("Please enter a score: ")
		_, err := fmt.Scanf("%d\n", &score)

		if err != nil {
			fmt.Println("Wrong input detected, please input a number")
			continue
		}

		if score >= 90 {
			fmt.Println("Selamat! Anda mendapatkan nilai A")
		} else if score < 90 && score >= 80 {
			fmt.Println("Anda mendapatkan nilai B")
		} else if score < 80 && score >= 70 {
			fmt.Println("Anda mendapatkan nilai C")
		} else if score < 70 && score >= 60 {
			fmt.Println("Anda mendapatkan nilai D")
		} else {
			fmt.Println("Anda mendapatkan nilai E")
		}
	}
}