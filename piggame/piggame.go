package piggame

import (
	"math/rand/v2"
)

func currentRoundScore(holdingScore int) int {

	currentScore := 0
	value := ((rand.IntN(5)) + 1)

	for value != 1 && currentScore < holdingScore {
		currentScore += value
		value = ((rand.IntN(5)) + 1)
	}

	if value == 1 {
		return 0
	}

	return currentScore
}

func playPigTwoPlayerPigGame(player1holdingScore int, player2holdingScore int) int {

	player1TotalScore, player2TotalScore := 0, 0
	currentTurn := 1

	for player1TotalScore < 100 && player2TotalScore < 100 {

		if currentTurn == 1 {
			player1TotalScore += currentRoundScore(player1holdingScore)
			currentTurn = 2
		} else {
			player2TotalScore += currentRoundScore(player2holdingScore)
			currentTurn = 1
		}
	}

	if player1TotalScore >= 100 {
		return 1
	} else {
		return 2
	}
}

func GetWinningProbability(iterations int, player1HoldingScore int, player2HoldingScore int) (float32, float32) {
	player1Wins, player2Wins := 0, 0
	for i := 0; i < iterations; i++ {
		result := playPigTwoPlayerPigGame(player1HoldingScore, player2HoldingScore)

		if result == 1 {
			player1Wins++
		} else {
			player2Wins++
		}
	}

	return (float32(player1Wins) / float32(iterations)), float32(player2Wins) / float32(iterations)
}
