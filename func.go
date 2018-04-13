package main

func testValues(expected, actual int) (bool, error) {
	return expected == actual, nil
}
