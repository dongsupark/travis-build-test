package main

import (
	"fmt"
	"os"
)

func main() {
	expected := 67
	actual := 67

	if expected != actual {
		fmt.Printf("mismatch! expected=%d, actual=%d\n", expected, actual)
		os.Exit(1)
	}

	fmt.Printf("match! expected=%d, actual=%d\n", expected, actual)
	os.Exit(0)
}
