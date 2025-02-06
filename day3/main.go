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
	fmt.Println(getMul(input))
}

func getMul(s string) []string {
	re := regexp.MustCompile("(?i)mul\\([A-Za-z0-9]+,[A-Za-z0-9]+\\)")
	return re.FindAllString(s, -1)
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
