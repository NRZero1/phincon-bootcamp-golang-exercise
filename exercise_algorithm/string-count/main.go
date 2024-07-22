package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"

	lower := strings.ToLower(sentence)
	
	resA := strings.Count(lower, "a")
	resI := strings.Count(lower, "i")
	resU := strings.Count(lower, "u")
	resE := strings.Count(lower, "e")
	resO := strings.Count(lower, "o")

	total := resA + resI + resU + resE + resO

	fmt.Println("Total vowel charactes is ", total)
}