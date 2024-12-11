package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	//"slices"
)

//const FILENAME = "example.txt"
const FILENAME = "input.txt"

func blink(stones []string) []string {
    newStones := make([]string, 0)
    for _, stone := range stones {
        s, _ := strconv.Atoi(stone)
        if s == 0 {
            newStones = append(newStones, "1")
        } else if len(stone) % 2 == 0 {
            lStone := strings.TrimLeft(stone[:len(stone)/2], "0")
            rStone := strings.TrimLeft(stone[len(stone)/2:], "0")
            //fmt.Println(stone, "Becomes", lStone, rStone)
            newStones = append(newStones, lStone)
            newStones = append(newStones, rStone)
        } else {
            newStones = append(newStones, strconv.Itoa(s * 2024))
        }
    }
    return newStones
}

func puzzle1(input string) {
    stones := strings.Split(input, " ")

    blinkTimes := 75
    for i := range blinkTimes {
        stones = blink(stones)
        // Split off and consolidate memory
        if len(stones) > 1000 {
        }
        fmt.Println(i, len(stones))
    }

    fmt.Println(len(stones))
}

func puzzle2(input string) {
}


func main() {
	dat, err := os.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	input := string(dat[0:len(dat)-1])

	puzzle1(input)
}

