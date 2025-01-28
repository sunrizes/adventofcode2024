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
	var total int = 0

	input := numberConversion(readFile("input.txt"))
	for _, v := range input {
		if checkSafety(v) {
			total++
		}
	}

	fmt.Println(total)
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

	if len(input) < 2 {
		fmt.Println("Not enough elements in slice, or slice is empty.")
		return false
	}

	if input[0] > input[1] {
		// Check decrease by 1, 2 or 3 and equality only

		for i := 0; i < len(input); i++ {
			if i+1 == len(input) {
				break
			}

			if input[i] < input[i+1] {
				fmt.Println("Number decreased after increasing")
				return false
			} else {
				if input[i]-input[i+1] == 1 || input[i]-input[i+1] == 2 || input[i]-input[i+1] == 3 {
					output = true
				} else {
					fmt.Println("Number decrease more than 3")
					return false
				}
			}
		}
	} else if input[0] < input[1] {
		// Check increase by 1, 2 or 3 and equality only

		for i := 0; i < len(input); i++ {
			if i+1 == len(input) {
				break
			}

			if input[i] > input[i+1] {
				fmt.Println("Number increased after decreasing")
				return false
			} else {

				if input[i+1]-input[i] == 1 || input[i+1]-input[i] == 2 || input[i+1]-input[i] == 3 {
					output = true
				} else {
					fmt.Println("Number increase more than 3")
					return false
				}

			}
		}
	} else {
		fmt.Println("No increase or decrease")
		return false
	}

	return output
}
