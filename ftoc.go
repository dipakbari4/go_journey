// This program demonstrates the use case
// of conversion method

package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g deg F = %g deg C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g deg F = %g deb C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
