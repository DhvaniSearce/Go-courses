// 3-Write a function with one variadic parameter that finds the greatest number in a list of numbers.
package main

import "fmt"

func max(num ...int) int {
	l := 0
	for _, n := range num {
		if n > l {
			l = n
		}
	}
	return l
}
func main() {
	maximum := max(12, 3, 4, 567, 3, 1, 2, 3, 0)
	fmt.Println(maximum)
}
