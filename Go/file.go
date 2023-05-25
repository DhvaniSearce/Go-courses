package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inpfile := os.Args[1]
	file, err := ioutil.ReadFile(inpfile)
	if err != nil {
		fmt.Printf("The file is not readable due to: %s \n", err)
	}
	fmt.Println(string(file))
}
