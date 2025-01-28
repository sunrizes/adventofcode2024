// Day 2 - Red-Nosed Reports

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := numberConversion(readFile("input.txt"))
	for _, v := range input {
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

	var output = []string{}

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func numberConversion(numbers []string) [][]int {
	var output = [][]int{}

	for _, v := range numbers {
		var tn = []int{}
		ts := strings.Fields(v)
		for i := 0; i < len(ts); i++ {
			n, _ := strconv.Atoi(ts[i])
			tn = append(tn, n)
		}
		output = append(output, tn)
	}

	return output
}
