package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile("(?i)mul\\([A-Za-z0-9]+,[A-Za-z0-9]+\\)")
	return re.FindAllString(s, -1)
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
