//even-odd

package main

func main() {
	numb := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numb {
		if n%2 == 0 {
			println(n, " is even")
		} else {
			println(n, " is odd")
		}
	}
}
