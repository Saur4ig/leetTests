package stringsE

import (
	"strconv"
)

// 412. Fizz Buzz
// Write a program that outputs the string representation of numbers from 1 to n.
// But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”.
// For numbers which are multiples of both three and five output “FizzBuzz”.
// link - https://leetcode.com/problems/fizz-buzz/
// bad solution
func fizzBuzz(n int) []string {
	resp := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			resp[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			resp[i-1] = "Fizz"
		} else if i%5 == 0 {
			resp[i-1] = "Buzz"
		} else {
			resp[i-1] = strconv.Itoa(i)
		}
	}
	return resp
}
