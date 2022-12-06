package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := os.Getenv("D1_PATH")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var caloriesOfOne int
	var caloriesOfAll []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row string

		for _, v := range scanner.Text() {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalf("not int: %v", err)
			}
			row = fmt.Sprintf("%v%v", row, i)
		}

		if row != "" {
			product, err := strconv.Atoi(row)
			if err != nil {
				log.Fatalf("cannot convert string to int: %v", err)
			}
			caloriesOfOne += product
		} else {
			caloriesOfAll = append(caloriesOfAll, caloriesOfOne)
			caloriesOfOne = 0
		}
	}
	fmt.Println(caloriesOfAll)
	max := findMax(caloriesOfAll)
	fmt.Printf("maximum calories: %v", max)
}

func findMax(input []int) int {
	var max int
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}
