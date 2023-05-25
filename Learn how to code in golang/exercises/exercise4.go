// 4-Create a program that prints to the terminal asking for a user to enter a small number and a larger number. Print the remainder of the larger number divided by the smaller number.
package main

import "fmt"

func main() {
	var min int
	var max int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&min)
	fmt.Print("Enter a large number: ")
	fmt.Scan(&max)
	x := max % min
	fmt.Printf("The remainder of %v divided by %v is equal to %v \n", max, min, x)
}
