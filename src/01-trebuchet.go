package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func removeAlphabeticCharacters(inputString string) string {
	re := regexp.MustCompile("[a-zA-Z]")
	return re.ReplaceAllString(inputString, "")
}

func getFirstAndLastDigits(digitLine string) (int, error) {
	if len(digitLine) > 0 {
		return strconv.Atoi((string)(digitLine[0]) + (string)(digitLine[len(digitLine)-1]))
	}
	return 0, nil
}

func getFirstDigit(digitLine string) (int, error) {
	if len(digitLine) > 0 {
		return strconv.Atoi((string)(digitLine[0]))
	}
	return 0, nil
}

func replacePatterns(input string, reverse bool) string {
	if reverse {
		input = reverseString(input)
	}
	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	output := input
	for i := 0; i < len(output); i++ {
		for pattern, replacement := range replacements {
			if reverse {
				pattern = reverseString(pattern)
			}
			patternLen := len(pattern)
			if i <= len(output)-patternLen && output[i:i+patternLen] == pattern {
				output = output[:i] + replacement + output[i+patternLen:]
				i += len(replacement) - 1
			}
		}
	}

	return output
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	fileContent := readFile("../inputs/01.in")

	calibrationLines := strings.Split(fileContent, "\n")

	// Part 1
	acc := 0

	for _, line := range calibrationLines {
		lineWithOnlyDigits := removeAlphabeticCharacters(line)
		calibrationValue, err := getFirstAndLastDigits(lineWithOnlyDigits)
		if err != nil {
			panic(err)
		}
		acc += calibrationValue
	}

	fmt.Println(acc)

	f1, err := os.Create("./test1")
	f2, err := os.Create("./test2")

	if err != nil {
		panic(err)
	}

	defer f1.Close()
	defer f2.Close()

	// Part 2
	acc2 := 0

	for _, line := range calibrationLines {
		fmt.Println(line)
		firstCharacterString := replacePatterns(line, false)
		lastCharacterString := replacePatterns(line, true)
		f1.WriteString(firstCharacterString + "\n")
		f2.WriteString(reverseString(lastCharacterString) + "\n")
		firstCharacterWithoutLetters := removeAlphabeticCharacters(firstCharacterString)
		lastCharacterWithoutLetters := removeAlphabeticCharacters(lastCharacterString)

		firstCalibrationValue, errFirst := getFirstDigit(firstCharacterWithoutLetters)
		lastCalibrationValue, errLast := getFirstDigit(lastCharacterWithoutLetters)
		calibrationValue := firstCalibrationValue + lastCalibrationValue
		fmt.Println(calibrationValue)
		if errFirst != nil || errLast != nil {
			panic(err)
		}
		acc2 += calibrationValue
	}

	fmt.Println(acc2)

}
