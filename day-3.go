package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3Solution(){
		file, err := os.Open("input-day-3.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		rows := make([]string, 0)
		for scanner.Scan() {
			line := scanner.Text()
			rows = append(rows, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(getRowParts(rows[0]))
			
			

}

func isByteDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func getRowParts(row string) []PossiblePart {
	rowParts := make([]PossiblePart, 0)
	index := 0
	for {
		if index >= len(row) {
			break
		}
		currentChar := row[index]
		partVal := string(currentChar)
		fmt.Println(currentChar)
		if isByteDigit(currentChar) {
			partStart := index
			for {
				index++
				if index >= len(row) {
					break
				}
				currentChar = row[index]
				if !isByteDigit(currentChar) {
					break
				}
				partVal += string(currentChar)
			}
			partEnd := index - 1
			partValInt, err := strconv.Atoi(partVal)
			if err != nil {
				fmt.Println(err)
			}
			rowParts = append(rowParts, PossiblePart{
				Start: partStart,
				End: partEnd,
				Value: partValInt,
			})
		} else {
			index++
		}
	}
	return rowParts
}

type PossiblePart struct {
	Start int
	End int
	Value int
}


func isValidPart(start int, end int, rows []string) bool {
	return true
}