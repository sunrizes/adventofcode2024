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
		fmt.Println(checkSafety(v))
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

func checkSafety(input []int) bool {
	var output bool = false

	// var temp int = 0

	if input[0] > input[1] {
		for i := 0; i < len(input); i++ {
			if i+1 == len(input) {
				break
			}

			if input[i] > input[i+1] {
				// Check decrease and equality only
			}
		}
	} else if input[0] < input[1] {
		for i := 0; i < len(input); i++ {
			if i+1 == len(input) {
				break
			}

			if input[i] < input[i+1] {
				// Check increase and equality only
			}
		}
	}

	fmt.Println()

	return output
}
