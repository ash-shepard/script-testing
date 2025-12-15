package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Compare Script")

	os.Setenv("SET_THIS_VARIABLE", "if you see this, it is good")

}
