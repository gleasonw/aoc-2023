package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameSet struct {
	Id int
	DrawSet []Draw
}

type Draw struct {
	Red int
	Blue int
	Green int
}

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func Day2Solution(){
	Day2Part2()
}

func Day2Part2() {
	games := parseInput()
	sumPowerSet := 0
	for _, game := range games {
		maxBlue := -1
		maxRed := -1
		maxGreen := -1
		for _, draw := range game.DrawSet {
			if draw.Blue > maxBlue {
				maxBlue = draw.Blue
			}
			if draw.Red > maxRed {
				maxRed = draw.Red
			}
			if draw.Green > maxGreen {
				maxGreen = draw.Green
			}
		}
		sumPowerSet += maxBlue * maxRed * maxGreen
	}
	fmt.Println(sumPowerSet)
}

func Day2Part1() {
	games := parseInput()

	validIdSum := 0
	for _, game := range games {
		if isValidGame(game) {
			validIdSum += game.Id
		}
	}
	fmt.Println(validIdSum)
}

func parseInput() []GameSet {
	file, err := os.Open("input-day-2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	games := make([]GameSet, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		gameData := parts[0]
		gameSetsString := parts[1]
		gameId := strings.Split(gameData, " ")[1]
		gameIdConverted, err := strconv.Atoi(gameId)
		if err != nil {
			fmt.Println(err)
		}

		gameSetsFinal := make([]Draw, 0)
		gameFinal := GameSet{
			Id: gameIdConverted,
			DrawSet: gameSetsFinal,
		}
		splitGameSets := strings.Split(gameSetsString, ";")
		for _, game := range splitGameSets {
			gameDraws := strings.Split(game, ",")
			draw := Draw{
				Red: 0,
				Blue: 0,
				Green: 0,
			}
			for _, currentDraw := range gameDraws {
				currentDraw = strings.TrimSpace(currentDraw)
				gameData := strings.Split(currentDraw, " ")
				color := gameData[1]
				count := gameData[0]
				countConverted, err := strconv.Atoi(count)
				if err != nil {
					fmt.Println(err)
				}
				if color == "red" {
					draw.Red = countConverted
				}
				if color == "blue" {
					draw.Blue = countConverted
				}
				if color == "green" {
					draw.Green = countConverted
				}
			}
			gameFinal.DrawSet = append(gameFinal.DrawSet, draw)
	}
		games = append(games, gameFinal)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return games
}

func isValidGame(game GameSet) bool {
	for _, draw := range game.DrawSet {
		if !isValidDraw(draw) {
			return false
		}
	}
	return true
}

func isValidDraw(draw Draw) bool {
	if draw.Red > maxRed {
		return false
	}
	if draw.Blue > maxBlue {
		return false
	}
	if draw.Green > maxGreen {
		return false
	}
	return true
}