package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func calculateDistance(totalTime, holdingTime int) int {
	remainingTime := totalTime - holdingTime
	return holdingTime * remainingTime
}

func findNumberOfWinningWay(time, distance int) int {
	acc := 0

	for i := 14; i < time; i++ {
		if calculatedDistance := calculateDistance(time, i); calculatedDistance > distance {
			acc++
		}
	}

	return acc
}

func parseStringArray(strings []string) (numbers []int) {
	for _, rawNumber := range strings {
		if rawNumber != "" {
			number, err := strconv.Atoi(rawNumber)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
	}
	return
}

func main() {
	rawInput := readFile("../inputs/06.in")

	inputs := strings.Split(rawInput, "\n")

	TimeLine := strings.Split(inputs[0], ":")
	rawJoinedTime := strings.ReplaceAll(strings.TrimSpace(TimeLine[1]), " ", "")
	rawTimes := strings.Split(strings.TrimSpace(TimeLine[1]), " ")
	times := parseStringArray(rawTimes)

	DistanceLine := strings.Split(inputs[1], ":")
	rawJoinedDistance := strings.ReplaceAll(strings.TrimSpace(DistanceLine[1]), " ", "")
	rawDistances := strings.Split(strings.TrimSpace(DistanceLine[1]), " ")
	distances := parseStringArray(rawDistances)

	acc := 1
	for i := 0; i < len(times); i++ {
		acc = acc * findNumberOfWinningWay(times[i], distances[i])
	}

	BigTime, timeErr := strconv.Atoi(rawJoinedTime)
	BigDistance, distErr := strconv.Atoi(rawJoinedDistance)

	if timeErr != nil || distErr != nil {
		panic(timeErr)
	}

	fmt.Printf("acc: %v\n", acc)
	fmt.Printf("findNumberOfWinningWay(BigTime, BigDistance): %v\n", findNumberOfWinningWay(BigTime, BigDistance))

}
