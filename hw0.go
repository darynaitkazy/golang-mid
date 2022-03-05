// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, à¤¦à¥à¤¨à¤¿à¤¯à¤¾!")
	fmt.Println(Fizzbuzz(15))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPalindrome("abac"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	var l int = 0
	var r int = len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
