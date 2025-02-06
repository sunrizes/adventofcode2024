package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	input := readFile("input.txt")
	fmt.Println(input)
}

func getMul(s string) bool {
	re := regexp.MustCompile("(?i)mul\\([A-Za-z0-9]+,[A-Za-z0-9]+\\)")
	return re.MatchString(s)
}

func readFile(fileName string) []string {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var output = []string{}

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}
