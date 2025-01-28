// Day 1 - Hystorian Hysteria

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	var ls = []int{}
	var rs = []int{}

	for _, v := range input {
		temp := strings.Fields(v)

		ln, _ := strconv.Atoi(temp[0])
		rn, _ := strconv.Atoi(temp[len(temp)-1])

		ls = append(ls, ln)
		rs = append(rs, rn)
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
