// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, Daryn!")
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(Anagram("abcd", "dcbb"))
	fmt.Println(FindEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SliceProduct([]int{3, 3, 4}))
	fmt.Println(Unique([]int{2, 2, 2, 3, 4, 2, 4, 6}))
	fmt.Println(InvertMap(map[string]int{
		"fir":  1,
		"sec":  2,
		"thir": 3,
	}))
	fmt.Println(TopCharacters("122222233333344444111111", 3))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var a [10]byte
	ind := 0
	for _, cur := range phone {
		if unicode.IsDigit(cur) {
			a[ind] = byte(cur)
			ind++
		}
	}
	var q string
	for i := 0; i <= 2; i++ {
		q += string(a[i])
	}
	var w string
	for i := 3; i <= 5; i++ {
		w += string(a[i])
	}
	var e string
	for i := 6; i <= 9; i++ {
		e += string(a[i])
	}
	ans := "(" + q + ")" + " " + w + "-" + e
	return ans
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	fir := strings.Split(s1, "")
	sort.Strings(fir)
	strings.Join(fir, "")

	sec := strings.Split(s2, "")
	sort.Strings(sec)
	strings.Join(sec, "")

	for i := 0; i < len(s1); i++ {
		if fir[i] != sec[i] {
			return false
		}
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var ans []int
	for _, i := range e {
		if i%2 == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	ans := 1
	for _, i := range e {
		ans *= i
	}
	return ans
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	mp := make(map[int]bool)
	for _, i := range e {
		mp[i] = true
	}
	var ans []int
	for i, _ := range mp {
		ans = append(ans, i)
	}
	return ans
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	ans := make(map[int]string)
	for i, j := range kv {
		ans[j] = i
	}
	return ans
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	ans := make(map[rune]int)
	for _, i := range s {
		ans[i]++
	}
	for i, j := range ans {
		if j <= k {
			delete(ans, i)
		}
	}
	return ans
}
