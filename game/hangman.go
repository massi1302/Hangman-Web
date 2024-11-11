package hangman

import (
	"hangman/game/state"
	"log"
	"math/rand"
	"unicode/utf8"
)

var gameStatePerUser = make(map[string]*state.GameState)

// NewGame initialise une nouvelle partie
func NewGame(username string, difficulty string) (*state.GameState, error) {
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

	if gameState, _ := state.LoadGameState(username); gameState != nil {
		gameStatePerUser[username] = gameState
		gameState.Difficulty = difficulty
		gameState.Word = *word
		gameState.DisplayedWord = displayedWord
		gameState.Lives = state.Lives
		gameState.UsedLetters = make([]string, 0)
		gameState.GameOver = false
	} else if gameState = gameStatePerUser[username]; gameState != nil {
		state.SaveGameState(username, gameState)
		gameState.Difficulty = difficulty
		gameState.Word = *word
		gameState.DisplayedWord = displayedWord
		gameState.Lives = state.Lives
		gameState.UsedLetters = make([]string, 0)
		gameState.GameOver = false
	} else {
		gameState = &state.GameState{
			Difficulty:    difficulty,
			Word:          *word,
			DisplayedWord: displayedWord,
			Lives:         state.Lives,
			UsedLetters:   make([]string, 0),
			GameOver:      false,
			Victory:       false,
			Score:         0,
		}
		gameStatePerUser[username] = gameState
	}

	return gameStatePerUser[username], nil
}

func Continue(username string) *state.GameState {
	gameState, err := state.LoadGameState(username)
	if err != nil {
		log.Printf("Error loading game state: %v\n", err)
	}
	gameStatePerUser[username] = gameState
	return gameStatePerUser[username]
}

func GetGameState(username string) *state.GameState {
	return gameStatePerUser[username]
}
