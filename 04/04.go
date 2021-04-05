package main

import (
	"fmt"
	"math"
)

func setPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	var even, odd, five, prime, primeSmall []int

	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			even = append(even, i)
		}
		if i%2 != 0 {
			odd = append(odd, i)
		}
		if i%5 == 0 {
			five = append(five, i)
		}
		if setPrime(i) {
			prime = append(prime, i)
		}
		if setPrime(i) && i < 100 {
			primeSmall = append(primeSmall, i)
		}
	}
	fmt.Println("even... ", even)
	fmt.Println("odd... ", odd)
	fmt.Println("five... ", five)
	fmt.Println("prime... ", prime)
	fmt.Println("prime less than 100... ", primeSmall)
}
