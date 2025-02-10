// Day 4: Ceres Search

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("input.txt")
	findXmas(input)
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

func findXmas(sa []string) {
	// Find X in line
	for i := 0; i < len(sa); i++ {
		line := sa[i]
		for j := 0; j < len(line); j++ {
			if line[j] == 'X' {
				fmt.Println("X Found at index", j, "in line", i+1)

				// Horizontal Search
				if j > 0 && j < len(line)-1 {
					// Search for M for X->M->A->S
					if line[j+1] == 'M' {
						fmt.Println("There's an M (right side) at index", j+1, "line", i+1)
					}
					if line[j-1] == 'M' {
						fmt.Println("There's an M (left side) at index", j-1, "line", i+1)
					}

					// Search for S for S->A->M->X
					if line[j+1] == 'S' {
						fmt.Println("There's an S (right side) at index", j+1, "line", i+1)
					}
					if line[j-1] == 'S' {
						fmt.Println("There's an S (left side) at index", j-1, "line", i+1)
					}
				}

				// Vertical Search

				// Diagonal Search
			}
		}
	}
}
