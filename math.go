package main

func CountDigit(x int) int {
	digitCount := 0
	for x > 0 {
		digitCount++
		x /= 10
	}
	return digitCount
}
