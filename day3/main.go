// Day 3 - Mull It Over

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")
	muls := getMul(input)

	var total int = 0

	for i := 0; i < len(muls); i++ {
		total += getProduct(getMultipliers(muls[i]))
	}

	fmt.Println(total)
}

func getMul(s string) []string {
	var output = []string{}

	ignore := false

	doReg := regexp.MustCompile(`do\(\)`)
	dontReg := regexp.MustCompile(`don't\(\)`)

	re := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if dontReg.MatchString(line) {
			ignore = true
		}

		if doReg.MatchString(line) {
			ignore = false
		}

		if !ignore {
			matches := re.FindAllString(line, -1)
			for _, match := range matches {
				output = append(output, match)
			}
		}
	}

	fmt.Println(output)

	return output
}

func getMultipliers(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}

func getProduct(sa []string) int {
	num1, err := strconv.Atoi(sa[0])
	if err != nil {
		fmt.Println(err)
	}

	num2, err := strconv.Atoi(sa[1])
	if err != nil {
		fmt.Println(err)
	}

	return num1 * num2
}

func readFile(fileName string) string {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var output string = ""

	for scanner.Scan() {
		output += scanner.Text()
	}

	return output
}
