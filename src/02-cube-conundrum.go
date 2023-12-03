package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	index int
	red   []int
	green []int
	blue  []int
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func parseColor(rawColor string) (red int, green int, blue int) {
	trimmedColor := strings.TrimSpace(rawColor)
	colorData := strings.Split(trimmedColor, " ")
	numberOfColor, error := strconv.Atoi(colorData[0])

	if error != nil {
		panic(error)
	}

	switch colorData[1] {
	case "red":
		red = numberOfColor
	case "green":
		green = numberOfColor
	case "blue":
		blue = numberOfColor
	}

	return
}

func (g *Game) parseGameSet(rawSets string) {
	splittedRawSets := strings.Split(rawSets, ";")
	for _, rawSet := range splittedRawSets {
		rawColors := strings.Split(rawSet, ",")
		for _, rawColor := range rawColors {
			red, green, blue := parseColor(rawColor)
			g.red = append(g.red, red)
			g.green = append(g.green, green)
			g.blue = append(g.blue, blue)
		}
	}
}

func (g *Game) getPower() int {
	return slices.Max(g.red) * slices.Max(g.green) * slices.Max(g.blue)
}

func parseGames(rawGames []string) []Game {
	games := []Game{}
	for i, rawGame := range rawGames {
		newGame := Game{}
		gameParts := strings.Split(rawGame, ":")
		newGame.index = i + 1
		newGame.parseGameSet(gameParts[1])
		games = append(games, newGame)
	}

	return games
}

func filterGames(games []Game, red int, green int, blue int) []Game {
	filteredGames := []Game{}
	for _, game := range games {
		if len(game.red) != 0 && slices.Max(game.red) <= red && len(game.green) != 0 && slices.Max(game.green) <= green && len(game.blue) != 0 && slices.Max(game.blue) <= blue {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}

func main() {
	fileContent := readFile("../inputs/02.in")

	rawGames := strings.Split(fileContent, "\n")

	games := parseGames(rawGames)

	// Part 1

	filteredGames := filterGames(games, 12, 13, 14)

	acc := 0

	for _, filteredGames := range filteredGames {
		acc += filteredGames.index
	}

	fmt.Printf("acc: %v\n", acc)

	// Part 2

	acc2 := 0

	for _, game := range games {
		acc2 += game.getPower()
	}

	fmt.Printf("acc2: %v\n", acc2)
}
