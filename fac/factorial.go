package fac

import (
	"fmt"
	"time"
)

var factVal int = 1
var i int = 1
var n int

func factorial(n int) int {

	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= int(i)
		}

	}
	return factVal
}

func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func Factorial() {
	// var n int
	// fmt.Print("Enter A number: ")
	// fmt.Scanln(&n)
	// var fac = factorial(n)
	// fmt.Println(n, fac)
	fmt.Print("Enter a positive integer between 0 - 50 : ")
	start := time.Now()
	fmt.Scan(&n)

	fmt.Println("Factorial is: ", factorial(n))
	fmt.Println(time.Since(start).Seconds())
	start2 := time.Now()
	fmt.Println("Factorial is: ", factorialRecursive(n))
	fmt.Println(time.Since(start2).Seconds())
}
