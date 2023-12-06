package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	leftPart  []int
	rightPart []int
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func parsePart(part []string) (parsedPart []int) {
	for _, numberInString := range part {
		number, err := strconv.Atoi(numberInString)
		if err != nil {
			panic(err)
		}
		parsedPart = append(parsedPart, number)
	}
	return
}

func parseScracthCards(rawCards []string) (parsedScratchCard []ScratchCard) {
	for _, rawCard := range rawCards {
		splittedParts := strings.Split(rawCard, ":")
		cardParts := strings.Split(splittedParts[1], "|")
		leftPart := strings.Split(strings.TrimSpace(cardParts[0]), " ")
		rightPart := strings.Split(strings.TrimSpace(cardParts[1]), " ")
		parsedScratchCard = append(parsedScratchCard, ScratchCard{leftPart: parsePart(leftPart), rightPart: parsePart(rightPart)})
	}

	return
}

func main() {
	rawInput := readFile("../inputs/04.in")

	rawScratchCards := strings.Split(rawInput, "\n")

	scratchCards := parseScracthCards(rawScratchCards)

	for _, scratchCard := range scratchCards {

	}
}
