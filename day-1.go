package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)



func day1Solution(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineSum := getLineTwoDigit(line)
        fmt.Println(lineSum)
        sum += lineSum
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    log.Println(sum)
}

var numStrings = map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "zero": 0,
}

var reversedNumStrings = map[string]int{
    "eno": 1,
    "owt": 2,
    "eerht": 3,
    "ruof": 4,
    "evif": 5,
    "xis": 6,
    "neves": 7,
    "thgie": 8,
    "enin": 9,
    "orez": 0,
}


func getLineTwoDigit(line string) int {
    firstDigit, lastDigit := getFirstDigit(line), getLastDigit(line)
    num, err := strconv.Atoi(firstDigit + lastDigit)
    if err != nil {
        log.Fatal(err)
    }
    
    
    return num
}

func getFirstDigit(line string) string{
    firstDigit := ""
    for i, r := range line {
        if unicode.IsDigit(r){
            firstDigit = string(r)
            return firstDigit
        }
        builtString := string(r)
        for j := i + 1; j < len(line) && len(builtString) < 6; j ++ {
            builtString += string(line[j])
            if num, ok := numStrings[builtString]; ok {
                firstDigit = strconv.Itoa(num);
                return firstDigit
            }
        }
    }
    return firstDigit
}

func getLastDigit(line string) string{
    lastDigit := ""
    for i := len(line) - 1; i >= 0; i -- {
        currentChar := line[i]
        if '0' <= currentChar && currentChar <= '9'{
            lastDigit = string(currentChar)
            return lastDigit
        }
        builtString := string(currentChar)
        for j := i - 1; j >= 0 && len(builtString) < 6; j -- {
            builtString += string(line[j])

            if num, ok := reversedNumStrings[builtString]; ok {
                lastDigit = strconv.Itoa(num);
                return lastDigit
            }
        }
    }
    return lastDigit
}



