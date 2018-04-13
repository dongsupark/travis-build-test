package main

import (
	"fmt"
	"os"
)

func test2() {
	expected := 27
	actual := 27

	res := testValues(expected, actual)
	if !res {
		fmt.Printf("mismatch! expected=%d, actual=%d\n", expected, actual)
		os.Exit(1)
	}

	fmt.Printf("match! expected=%d, actual=%d\n", expected, actual)
}
