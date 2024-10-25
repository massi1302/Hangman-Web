package hangman

import (
	"encoding/json"
	"math/rand"
	"strings"
)

// GameState représente l'état actuel d'une partie
type GameState struct {
	Word          string   `json:"word"`
	DisplayedWord []string `json:"displayedWord"`
	Lives         int      `json:"lives"`
	UsedLetters   []string `json:"usedLetters"`
	GameOver      bool     `json:"gameOver"`
	Victory       bool     `json:"victory"`
	Score         int      `json:"score"`
}

// NewGame initialise une nouvelle partie
func NewGame(difficulty string) *GameState {
	// Choisir le fichier de mots selon la difficulté
	filename := "words.txt"
	switch difficulty {
	case "EASY":
		filename = "easy.txt"
	case "MEDIUM":
		filename = "medium.txt"
	case "HARD":
		filename = "hard.txt"
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

	return &GameState{
		Word:          word,
		DisplayedWord: displayedWord,
		Lives:         10,
		UsedLetters:   make([]string, 0),
		GameOver:      false,
		Victory:       false,
		Score:         0,
	}
}

// GuessLetter traite une tentative de lettre
func (g *GameState) GuessLetter(letter string) bool {
	letter = strings.ToLower(letter)

	// Vérifier si la lettre a déjà été utilisée
	for _, usedLetter := range g.UsedLetters {
		if usedLetter == letter {
			return false
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
	}

	// Vérifier si la partie est terminée
	g.checkGameEnd()

	return correct
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
