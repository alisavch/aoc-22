package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	path := os.Getenv("D3_PATH")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var priority int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string
		line = scanner.Text()
		recurring := findEquals(line)
		if recurring != "" {
			priority += getPriority([]byte(recurring))
		}
	}

	fmt.Printf("the sum of the priorities: %v", priority)
}

func getPriority(str []byte) int {
	lowerCase := 1
	upperCase := 27
	symbol := str[0]
	if symbol >= 97 && symbol <= 122 {
		for s := 'a'; s <= 'z'; s++ {
			if string(str) == string(s) {
				return lowerCase
			}
			lowerCase++
		}
	} else if symbol >= 65 && symbol <= 90 {
		for s := 'A'; s <= 'Z'; s++ {
			if string(str) == string(s) {
				return upperCase
			}
			upperCase++
		}
	}
	return 0
}

func findEquals(str string) string {
	strLen := utf8.RuneCount([]byte(str))
	part1 := str[:strLen/2]
	part2 := str[strLen/2:]
	for _, v := range part1 {
		for _, vv := range part2 {
			if v == vv {
				return string(v)
			}
		}
	}
	return ""
}
