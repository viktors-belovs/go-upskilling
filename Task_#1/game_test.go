package main

import (
	"strings"
	"testing"
)

func TestIdentifyWinner(t *testing.T) {
	tests := []struct {
		playerChoice    string
		computerChoice  string
		expectedOutcome string
	}{
		{"paper", "paper", "Tie!"},
		{"paper", "stone", "You win!"},
		{"paper", "scissors", "Computer wins!"},
		{"stone", "paper", "Computer wins!"},
		{"stone", "stone", "Tie!"},
		{"stone", "scissors", "You win!"},
		{"scissors", "paper", "You win!"},
		{"scissors", "stone", "Computer wins!"},
		{"scissors", "scissors", "Tie!"},
	}

	for _, test := range tests {
		t.Run(test.playerChoice+" vs "+test.computerChoice, func(t *testing.T) {
			actual := identifyWinner(test.playerChoice, test.computerChoice)
			if !strings.Contains(actual, test.expectedOutcome) {
				t.Errorf("Expected outcome '%s', but got '%s'.", test.expectedOutcome, actual)
			}
		})
	}
}
