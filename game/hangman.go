package game

import (
	"hangman/game/state"
	"log"
	"math/rand"
	"unicode/utf8"
)

// /assets/images/hangman-9.png
var hangmanDraw = make(map[int]string)

var gameStatePerUser = make(map[string]*state.GameState)

const lives = 10

func init() {
	hangmanDraw[1] = "/assets/images/hangman-0.png"
	hangmanDraw[2] = "/assets/images/hangman-1.png"
	hangmanDraw[3] = "/assets/images/hangman-2.png"
	hangmanDraw[4] = "/assets/images/hangman-3.png"
	hangmanDraw[5] = "/assets/images/hangman-4.png"
	hangmanDraw[6] = "/assets/images/hangman-5.png"
	hangmanDraw[7] = "/assets/images/hangman-6.png"
	hangmanDraw[8] = "/assets/images/hangman-7.png"
	hangmanDraw[9] = "/assets/images/hangman-8.png"
	hangmanDraw[10] = "/assets/images/hangman-9.png"
}

// NewGame initialise une nouvelle partie
func NewGame(user string, difficulty string) (*state.GameState, error) {
	// Choisir le fichier de mots selon la difficulté
	var path string
	switch difficulty {
	case "EASY":
		path = "./resources/data/easy_words.txt"
	case "MEDIUM":
		path = "./resources/data/medium_words.txt"
	case "HARD":
		path = "./resources/data/hard_words.txt"
	default:
		path = "./resources/data/easy_words.txt"
	}

	word, err := RandomWord(path)
	if err != nil {
		log.Printf("Error generating random word: %v\n", err)
		return nil, err
	}

	displayedWord := make([]string, utf8.RuneCountInString(*word))
	for i := range displayedWord {
		displayedWord[i] = "_"
	}

	// Révéler quelques lettres au début selon la difficulté
	nbLettersToReveal := 0
	switch difficulty {
	case "EASY":
		nbLettersToReveal = 2
	case "MEDIUM":
		nbLettersToReveal = 1
	}

	// Révéler des lettres aléatoires
	revealed := make(map[int]bool)
	for i := 0; i < nbLettersToReveal; i++ {
		randIndex := rand.Intn(utf8.RuneCountInString(*word))
		for revealed[randIndex] {
			randIndex = rand.Intn(utf8.RuneCountInString(*word))
		}
		revealed[randIndex] = true
		displayedWord[randIndex] = string([]rune(*word)[randIndex])
	}

	gameState := gameStatePerUser[user]

	if gameState != nil {
		gameState.Difficulty = difficulty
		gameState.Word = *word
		gameState.DisplayedWord = displayedWord
		gameState.Lives = lives
		gameState.UsedLetters = make([]string, 0)
		gameState.GameOver = false
	} else {
		gameState = &state.GameState{
			Difficulty:    difficulty,
			Word:          *word,
			DisplayedWord: displayedWord,
			Lives:         lives,
			UsedLetters:   make([]string, 0),
			GameOver:      false,
			Victory:       false,
			Score:         0,
		}
		gameStatePerUser[user] = gameState
	}

	return gameState, nil
}

func Continue(username string) *state.GameState {
	gameState, err := state.LoadGameState(username)
	if err != nil {
		log.Printf("Error loading game state: %v\n", err)
	}
	gameStatePerUser[username] = gameState
	return gameStatePerUser[username]
}

func GetGameState(user string) *state.GameState {
	return gameStatePerUser[user]
}
