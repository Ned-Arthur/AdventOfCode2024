package main

import (
	"fmt"
	"os"
)

const FILENAME = "example.txt"

func puzzle1(input string) {
	fmt.Println(input)
}

func main() {
	dat, err := os.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	input := string(dat[0:len(dat)-1])

	puzzle1(input)
}

