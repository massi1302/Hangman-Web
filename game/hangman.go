package hangman

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"strings"
)

var gameStatePerUser = make(map[string]*GameState)

// GameState représente l'état actuel d'une partie
type GameState struct {
	Word          string
	DisplayedWord []string
	Lives         int
	UsedLetters   []string
	UnusedLetters []string
	GameOver      bool
	Victory       bool
	Score         int
	Hearts        int
}

type FilteredGameState struct {
	DisplayedWord []string
	Lives         int
	UsedLetters   []string
	UnusedLetters []string
	GameOver      bool
	Victory       bool
	Score         int
	Hearts        int
}

// NewGame initialise une nouvelle partie
func NewGame(user string, difficulty string) *FilteredGameState {
	// Choisir le fichier de mots selon la difficulté
	filename := "words.txt"
	switch difficulty {
	case "EASY":
		filename = "./data/easy_words.txt"
	case "MEDIUM":
		filename = "./data/medium_words.txt"
	case "HARD":
		filename = "./data/hard_words.txt"
	}

	word := RandomWord(filename)
	displayedWord := make([]string, len(word))
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
		randIndex := rand.Intn(len(word))
		for revealed[randIndex] {
			randIndex = rand.Intn(len(word))
		}
		revealed[randIndex] = true
		displayedWord[randIndex] = string(word[randIndex])
	}

	if gameStatePerUser[user] != nil {
		gameStatePerUser[user].Word = word
		gameStatePerUser[user].DisplayedWord = displayedWord
		gameStatePerUser[user].Lives = 6
		gameStatePerUser[user].UsedLetters = make([]string, 0)
		gameStatePerUser[user].UnusedLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		gameStatePerUser[user].GameOver = false
		gameStatePerUser[user].Hearts = 6
	} else {
		gameStatePerUser[user] = &GameState{
			Word:          word,
			DisplayedWord: displayedWord,
			Lives:         6,
			UsedLetters:   make([]string, 0),
			UnusedLetters: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
			GameOver:      false,
			Victory:       false,
			Score:         0,
			Hearts:        6,
		}
	}

	return &FilteredGameState{
		DisplayedWord: gameStatePerUser[user].DisplayedWord,
		Lives:         gameStatePerUser[user].Lives,
		UsedLetters:   gameStatePerUser[user].UsedLetters,
		UnusedLetters: gameStatePerUser[user].UnusedLetters,
		GameOver:      gameStatePerUser[user].GameOver,
		Victory:       gameStatePerUser[user].Victory,
		Score:         gameStatePerUser[user].Score,
		Hearts:        gameStatePerUser[user].Hearts,
	}
}

func (g *GameState) ToFilteredGameState() *FilteredGameState {
	return &FilteredGameState{
		DisplayedWord: g.DisplayedWord,
		Lives:         g.Lives,
		UsedLetters:   g.UsedLetters,
		UnusedLetters: g.UnusedLetters,
		GameOver:      g.GameOver,
		Victory:       g.Victory,
		Score:         g.Score,
		Hearts:        g.Hearts,
	}
}

func GetGameState(user string) *GameState {
	return gameStatePerUser[user]
}

// GuessLetter traite une tentative de lettre
func (g *GameState) GuessLetter(letter string) *GameState {
	letter = strings.ToLower(letter)

	// Vérifier si la lettre a déjà été utilisée
	for _, usedLetter := range g.UsedLetters {
		if usedLetter == letter {
			return g
		}
	}

	// Ajouter la lettre aux lettres utilisées
	g.UsedLetters = append(g.UsedLetters, letter)

	// Vérifier si la lettre est dans le mot
	correct := false
	for i, char := range g.Word {
		if string(char) == letter {
			g.DisplayedWord[i] = letter
			correct = true
		}
	}

	// Si la lettre n'est pas dans le mot, perdre une vie
	if !correct {
		g.Lives--
		g.Hearts--
	}

	// Vérifier si la partie est terminée
	g.checkGameEnd()

	return g
}

// GuessWord traite une tentative de mot complet
func (g *GameState) GuessWord(word string) bool {
	word = strings.ToLower(word)
	if word == strings.ToLower(g.Word) {
		for i, char := range g.Word {
			g.DisplayedWord[i] = string(char)
		}
		g.Victory = true
		g.GameOver = true
		g.Score = g.Lives * 100
		return true
	}
	g.Lives -= 2
	g.checkGameEnd()
	return false
}

// checkGameEnd vérifie si la partie est terminée
func (g *GameState) checkGameEnd() {
	// Vérifier la victoire
	if strings.Join(g.DisplayedWord, "") == strings.ToLower(g.Word) {
		g.Victory = true
		g.GameOver = true
		g.Score = g.Lives * 100
		return
	}

	// Vérifier la défaite
	if g.Lives <= 0 {
		g.GameOver = true
		g.Victory = false
		g.Score = 0
	}
}

// ToJSON convertit l'état du jeu en JSON
func (g *GameState) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}

func (f *FilteredGameState) DrawHearts() template.HTML {
	heartIcon := "❤️"
	var heartString string
	for i := 0; i < f.Hearts; i++ {
		heartString += heartIcon + " "
	}
	return template.HTML(heartString)
}
