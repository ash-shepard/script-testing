package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Attempting to get a variable passed from github actions")

	value, exists := os.LookupEnv("GET_THIS_VARIABLE")
	
	if exists {
		fmt.Println("Value of GET_THIS_VARIABLE:", value)
	} else {
		fmt.Println("GET_THIS_VARIABLE is not set")
	}

}
