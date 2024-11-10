package util

import hangman "hangman/game"

type GameData struct {
	FilteredGameState *hangman.FilteredGameState
	Alphabet          []string
}

func NewGameData(filteredGameState *hangman.FilteredGameState) *GameData {
	return &GameData{
		FilteredGameState: filteredGameState,
		Alphabet:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	}
}
