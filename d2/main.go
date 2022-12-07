package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	LOSS     = 0
	DRAW     = 3
	VICTORY  = 6
)

func main() {
	path := os.Getenv("D2_PATH")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var games []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	var totalScore int
	for _, game := range games {
		totalScore += calculatePerRound(game)
	}

	fmt.Printf("total score: %v", totalScore)

}

func calculatePerRound(game string) int {
	participants := map[string]int{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSORS,
		"Z": SCISSORS,
	}
	p1 := participants[string(game[0])]
	p2 := participants[string(game[2])]
	result := calculate(p1, p2)
	return result + p2
}

func calculate(player1, player2 int) int {
	if player2 == player1 {
		return DRAW
	} else if player2 == ROCK && player1 == SCISSORS {
		return VICTORY
	} else if player2 == SCISSORS && player1 == PAPER {
		return VICTORY
	} else if player2 == PAPER && player1 == ROCK {
		return VICTORY
	} else {
		return LOSS
	}
}
