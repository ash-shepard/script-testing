package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Attempting to set a variable to pass back to github actions")

	os.Setenv("SET_THIS_VARIABLE", "if you see this, it is good")

}
