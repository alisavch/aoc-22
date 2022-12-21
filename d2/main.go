package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Move int

const (
	ROCK Move = iota + 1
	PAPER
	SCISSORS
)

type Result int

const (
	LOSS Result = iota // LOSS = 0
	_
	_
	DRAW // DRAW =3
	_
	_
	VICTORY // VICTORY = 6
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

	var totalScore Result
	for _, game := range games {
		totalScore += calculatePerRound(game)
	}

	fmt.Printf("total score: %v", totalScore)

}

func calculatePerRound(game string) Result {
	participants := map[string]Move{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSORS,
		"Z": SCISSORS,
	}
	p1 := participants[string(game[0])]
	p2 := participants[string(game[2])]
	roundResult := calculate(p1, p2)
	finalScore := int(roundResult) + int(p2)
	return Result(finalScore)
}

func calculate(player1, player2 Move) Result {
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
