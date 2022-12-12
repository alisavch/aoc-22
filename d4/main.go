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
	path := os.Getenv("D4_PATH")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var totalScore int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contains := contentCheck(scanner.Text())
		if contains {
			totalScore++
		}
	}
	fmt.Printf("number of assignment pairs one range contains completely: %v", totalScore)
}

func contentCheck(row string) bool {
	pairsOFNumbers := split(row)
	return include(pairsOFNumbers)
}

func split(row string) [][]int {
	var pairs [][]int
	arrays := strings.Split(row, ",")
	for _, v := range arrays {
		values := strings.Split(v, "-")
		var pair []int
		for _, vv := range values {
			i, err := strconv.Atoi(vv)
			if err != nil {
				log.Fatalf("cannot convert string to int")
			}
			pair = append(pair, i)
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

func include(pairs [][]int) bool {
	a, b := pairs[0][0], pairs[0][1]
	x, y := pairs[1][0], pairs[1][1]
	if a >= x && a <= y && b >= x && b <= y {
		return true
	} else if x >= a && x <= b && y >= a && y <= b {
		return true
	}
	return false
}
