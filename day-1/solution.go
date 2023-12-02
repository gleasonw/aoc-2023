package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineSum := getLineTwoDigit(line)
        sum += lineSum
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    log.Println(sum)
}

func getLineTwoDigit(line string) int {
    first, last := "", ""
    for _, r := range line {
        if unicode.IsDigit(r) {
            if first == "" {
                first = string(r)
            } else {
                last = string(r)
            }
        }
    }
    if last == "" {
        last = first
    }
    num, err := strconv.Atoi(first + last)
    if err != nil {
        log.Fatal(err)
    }
    return num
}
