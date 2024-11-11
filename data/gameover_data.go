package data

import (
	"hangman/game/state"
)

type GameOverData struct {
	FilteredGameState *state.FilteredGameState
	Alphabet          []string
	Word              string
}

func NewGameOverData(gameState *state.GameState) *GameOverData {
	return &GameOverData{
		FilteredGameState: gameState.ToFilteredGameState(),
		Alphabet:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
		Word:              gameState.Word,
	}
}
