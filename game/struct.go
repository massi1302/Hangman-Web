package hangman

import (
	"time"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type PageData struct {
	Title          string
	Game           Game
	Message        string
	MessageType    string
	Scores         []Score
	CurrentSession string
}

type Game struct {
	GuessedWord    []string
	PlayerName     string
	GuessedLetters map[rune]string
	Diff           Difficulty
	MaxTries       int
	Word           *Word
}

type Word struct {
	revealedWord  []rune
	revealedCount int
	Value         string
}

type Score struct {
	PlayerName string    `json:"playerName"`
	Score      int       `json:"score"`
	Word       string    `json:"word"`
	Difficulty string    `json:"difficulty"`
	Date       time.Time `json:"date"`
}
