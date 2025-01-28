// Day 1 - Hystorian Hysteria

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	fmt.Println(sortAndCompare(ls, rs))
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

func sortAndCompare(ls []int, rs []int) int {
	slices.Sort(ls)
	slices.Sort(rs)

	var total int = 0
	var distances = []int{}

	for i := 0; i < len(ls); i++ {
		if ls[i] >= rs[i] {
			distances = append(distances, ls[i]-rs[i])
		} else {
			distances = append(distances, rs[i]-ls[i])
		}
	}

	for _, v := range distances {
		total += v
	}

	return total
}
