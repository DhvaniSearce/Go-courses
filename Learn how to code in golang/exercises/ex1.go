// 1-Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. Thesecond return should be a bool that let’s us know whether or not the argument was even. For example:
package main

import (
	"fmt"
)

func checker(n int) (int, bool) {
	return n / 2, n%2 == 0
}
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	x, y := checker(num)
	fmt.Printf("(%v,%v)\n", x, y)
}
