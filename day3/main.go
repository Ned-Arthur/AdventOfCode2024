package main

import (
	"fmt"
	"os"
	"strconv"
)

const FILENAME = "input.txt"

func c(b byte) string {
	return string(b)
}

func puzzle1(input string) {
	total := 0
	for i := range input {
		if len(input) - i >= 4 && input[i:i+4] == "mul(" {
			j := i+4
			num1 := ""
			for c(input[j]) != "," {
				num1 += c(input[j])
				j++
			}
			j++
			num2 := ""
			for c(input[j]) != ")" {
				num2 += c(input[j])
				j++
			}

			x1, err1 := strconv.Atoi(num1)
			x2, err2 := strconv.Atoi(num2)
			if err1 != nil || err2 != nil {
				continue
			}

			sum := x1 * x2
			total += sum
		}
	}
	fmt.Println(total)
}

func puzzle2(input string) {
	total := 0
	do := true
	for i := range input {
		if len(input) - i >= 7 && input[i:i+7] == "don't()" {
			do = false
			continue
		}
		if len(input) - i >= 4 && input[i:i+4] == "do()" {
			do = true
			continue
		}
		if do && len(input) - i >= 4 && input[i:i+4] == "mul(" {
			j := i+4
			num1 := ""
			for c(input[j]) != "," {
				num1 += c(input[j])
				j++
			}
			j++
			num2 := ""
			for c(input[j]) != ")" {
				num2 += c(input[j])
				j++
			}

			x1, err1 := strconv.Atoi(num1)
			x2, err2 := strconv.Atoi(num2)
			if err1 != nil || err2 != nil {
				continue
			}

			sum := x1 * x2
			total += sum
		}
	}
	fmt.Println(total)
}

func main() {
	dat, err := os.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	input := string(dat[0:len(dat)-1])

	puzzle2(input)
}

