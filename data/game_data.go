package data

import (
	"hangman/game/state"
)

type ScoresData struct {
	Username   string
	GameStates []*state.GameState
}

type GameData struct {
	FilteredGameState *state.FilteredGameState
	Alphabet          []string
}

func NewGameData(gameState *state.GameState) *GameData {
	return &GameData{
		FilteredGameState: gameState.ToFilteredGameState(),
		Alphabet:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	}
}

func NewScoresData(username string, gameStates []*state.GameState) *ScoresData {
	return &ScoresData{
		Username:   username,
		GameStates: gameStates,
	}
}
