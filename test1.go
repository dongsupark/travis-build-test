package main

import (
	"fmt"
	"os"
)

func test1() {
	expected := 17
	actual := 17

	res, err := testValues(expected, actual)
	if !res || err != nil {
		fmt.Printf("mismatch! expected=%d, actual=%d\n", expected, actual)
		os.Exit(1)
	}

	fmt.Printf("match! expected=%d, actual=%d\n", expected, actual)
}
