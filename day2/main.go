package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

const FILENAME = "input.txt"

func puzzle1(input string) {
    lines := strings.Split(input, "\n")

    numSafe := 0

    for _, line := range lines {
        stringReports := strings.Split(line, " ")
        reports := make([]int, 0)
        for _, report := range stringReports {
            rep, _ := strconv.Atoi(report)
            reports = append(reports, rep)
        }

        safe := checkReports(reports)

        if safe {
            numSafe++
        }
    }

    fmt.Println(numSafe)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func checkReports(reports []int) bool {
    isDescending := (reports[0] - reports[1]) > 0

    for i, report := range reports {
        if i == 0 {
            continue
        }
        diff := reports[i-1] - report
        if diff > 0 != isDescending || diff == 0 || abs(diff) > 3 {
            return false
        }
    }
    // All good, let the program know
    return true
}

func removeItem(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func puzzle2(input string) {
    lines := strings.Split(input, "\n")

    numSafe := 0

    for _, line := range lines {
        stringReports := strings.Split(line, " ")
        reports := make([]int, 0)
        for _, report := range stringReports {
            rep, _ := strconv.Atoi(report)
            reports = append(reports, rep)
        }

        safe := checkReports(reports)
        if !safe {
            for i:=0; i < len(reports); i++  {
                newReports := removeItem(reports, i)
                if checkReports(newReports) == true {
                    safe = true
                }
            }
        }

        if safe {
            numSafe++
        }
    }

    fmt.Println(numSafe)
}

func main() {
    dat, err := os.ReadFile(FILENAME)
    if err != nil {
        panic(err)
    }
    input := string(dat[0:len(dat)-1])

    puzzle2(input)
}

