package main

import (
	"fmt"
	"math"
)

func main() {
	left, right := 10, 19
	primes, closestPair := closetPrimeNumber(left, right)
	fmt.Println("All Prime Numbers:", primes)
	fmt.Println("Closest Prime Pair:", closestPair)
}

func closetPrimeNumber(left, right int) ([]int, []int) {
	var primeList []int
	// Function to find the closest prime numbers in range [left, right]
	for i := left; i <= right; i++ {
		if isPrime(i) {
			primeList = append(primeList, i)
		}
	}
	// Step 2: If less than 2 primes, return [-1, -1]
	if len(primeList) < 2 {
		return primeList, []int{-1, -1}
	}
	// Step 3: Find the closest prime pair
	minDiff := math.MaxInt32
	result := []int{-1, -1}
	for i := 1; i < len(primeList); i++ {
		diff := primeList[i] - primeList[i-1]
		if diff < minDiff {
			minDiff = diff
			result = []int{primeList[i-1], primeList[i]}
		}
	}
	return primeList, result
}

func isPrime(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
