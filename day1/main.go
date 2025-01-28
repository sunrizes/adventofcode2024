// Day 1 - Hystorian Hysteria

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	temp := readFile("input.txt")

	for _, v := range temp {
		fmt.Println(v)
	}
}

func readFile(fileName string) []string {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var temp = []string{}

	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}

	return temp
}
