package main

import (
	"math"
	"strconv"
)

// checks if a number is a prime else return false
//
// number divisible by 1 and itself only
func isPrime(n int) bool {
	// smallest prime number is 2
	if n == 2 {
		return true
	}
	// all even numbers except 2 aren't prime numbers
	if n < 2 || n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// number which it's sum is equall to the sum of it's positive divisors
// excluding the number itself
func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// not divisible by 2 will return true
func isOdd(n int) bool {
	return n%2 != 0
}

// equal to the sum of it's digits, raised to a power equal to it's length
//
// 153 -> 1^3 + 5^3 + 3^3 (len(153) = 3)
func isArmstrong(n int) bool {
	if n < 0 {
		return true
	}
	// convert number to a string, then create a slice containing the numbers
	numSlice := make([]int, 0, 6)
	numString := strconv.Itoa(n)
	for i := range numString {
		// convert to int so as to append to the slice
		numInt, _ := strconv.Atoi(string(numString[i]))
		numSlice = append(numSlice, numInt)
	}
	// now we have an array of numbers where we can index them
	// find the cube power, use an array to store the cube for each val
	lenNum := len(numSlice)
	numCubes := make([]int, 0, lenNum)
	for i := 0; i <= lenNum-1; i++ {
		x3 := int(math.Pow(float64(numSlice[i]), float64(lenNum)))
		numCubes = append(numCubes, x3)
	}
	// add sum of the cubes
	sum := 0
	for _, v := range numCubes {
		sum += v
	}
	// compare sum to the inital number
	return sum == n
}

// sum of all the digits e.g 11 -> 2
func sumDigits(n int) int {
	sum := 0
	math.Abs(float64(n))
	for n != 0 {
		// get the last digit -> 421 will get 1, 2 and finally 4
		sum += n % 10
		// remove the last digit
		n /= 10
	}
	return sum
}
