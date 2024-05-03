package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		fmt.Println("Please choose 'paper', 'stone' or 'scissors':")
		var playerChoice string
		fmt.Scanln(&playerChoice)

		computerChoice := getComputerChoice()

		fmt.Printf("You chose %s, computer chose %s\n", playerChoice, computerChoice)

		result := identifyWinner(playerChoice, computerChoice)
		fmt.Println(result)

		fmt.Println("Would you like to play again? (y/n)")
		var playAgain string
		fmt.Scanln(&playAgain)

		if playAgain != "y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func getComputerChoice() string {
	choices := []string{"paper", "stone", "scissors"}
	randomChoice := rand.Intn(len(choices))
	return choices[randomChoice]
}

func identifyWinner(playerChoice, computerChoice string) string {
	switch {
	case playerChoice == computerChoice:
		return "Tie!"
	case playerChoice == "paper" && computerChoice == "stone":
		return "You win!"
	case playerChoice == "stone" && computerChoice == "scissors":
		return "You win!"
	case playerChoice == "scissors" && computerChoice == "paper":
		return "You win!"
	default:
		return "Computer wins!"
	}
}
