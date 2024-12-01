package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
)

const FILENAME = "input.txt"

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func puzzle1(input string) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		var item1, item2 int
		_, err := fmt.Sscanf(line, "%d   %d", &item1, &item2)
		if err != nil { panic(err) }
		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDiff := 0

	for i := range list1 {
		abs(list1[i] - list2[i])
		totalDiff += abs(list1[i] - list2[i])
	}

	fmt.Println(totalDiff)
}

func countInList(num int, list []int) int {
	count := 0
	for _, element := range list {
		if element == num {
			count++
		}
	}
	return count
}

func puzzle2(input string) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		var item1, item2 int
		_, err := fmt.Sscanf(line, "%d   %d", &item1, &item2)
		if err != nil { panic(err) }
		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}

	similarity := 0

	for _, item1 := range list1 {
		count := countInList(item1, list2)
		similarity += item1 * count
	}

	fmt.Println(similarity)
}

func main() {
	dat, err := os.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	input := string(dat[0:len(dat)-1])

	puzzle2(input)
}

