package main

import (
	"fmt"
	"go-bootcamp/piggame"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide the holding score of the players")
	}

	player1HoldingScore, _ := strconv.Atoi(args[0])
	player2HoldingScore, _ := strconv.Atoi(args[1])

	player1WinPercent, player2WinPercent := piggame.GetWinningProbability(200, player1HoldingScore, player2HoldingScore)

	fmt.Println("Win percent returned is ", player1WinPercent, player2WinPercent)

	fmt.Printf("Holding at %d vs Holding at %d: wins: %.2f%%, losses: %.2f%%\n", player1HoldingScore, player2HoldingScore, (player1WinPercent * 100), (player2WinPercent * 100))
}
