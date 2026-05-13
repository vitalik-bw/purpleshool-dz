package main

import "fmt"

func main() {
	const usdEUR = 0.94
	const usdRUB = 79
	eurRUB := usdRUB / usdEUR
	fmt.Print(eurRUB)
}

func inputUser() int {
	var input int
	fmt.Scan(&input)
	return input
}

func calculateExchange(amount float64, fromCurrency string, toCurrency string) float64 {
	return amount
}
