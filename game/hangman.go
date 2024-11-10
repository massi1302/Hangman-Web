package hangman

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"strings"
	"unicode/utf8"
)

// /assets/images/hangman-9.png
var hangmanDraw = make(map[int]string)

var gameStatePerUser = make(map[string]*GameState)

const lives = 10

// GameState représente l'état actuel d'une partie
type GameState struct {
	Difficulty    string
	Word          string
	DisplayedWord []string
	Lives         int
	UsedLetters   []string
	GameOver      bool
	Victory       bool
	Score         int
}

type FilteredGameState struct {
	Difficulty    string
	DisplayedWord []string
	Lives         int
	UsedLetters   []string
	GameOver      bool
	Victory       bool
	Score         int
}

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
func NewGame(user string, difficulty string) (*FilteredGameState, error) {
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
		gameState = &GameState{
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

	return gameState.ToFilteredGameState(), nil
}

func (g *GameState) ToFilteredGameState() *FilteredGameState {
	return &FilteredGameState{
		Difficulty:    g.Difficulty,
		DisplayedWord: g.DisplayedWord,
		Lives:         g.Lives,
		UsedLetters:   g.UsedLetters,
		GameOver:      g.GameOver,
		Victory:       g.Victory,
		Score:         g.Score,
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
func (g *GameState) checkGameEnd() bool {
	// Vérifier la victoire
	if strings.Join(g.DisplayedWord, "") == strings.ToLower(g.Word) {
		g.Victory = true
		g.GameOver = true
		g.Score += g.Lives * 100
		return true
	}

	// Vérifier la défaite
	if g.Lives <= 0 {
		g.GameOver = true
		g.Victory = false
		return true
	}

	return false
}

// ToJSON convertit l'état du jeu en JSON
func (g *GameState) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}

func (f *FilteredGameState) DrawHearts() template.HTML {
	heartIcon := "❤️"
	var heartString string
	for i := 0; i < f.Lives; i++ {
		heartString += heartIcon + " "
	}
	return template.HTML(heartString)
}

func (f *FilteredGameState) HangmanDraw(remainingLives int) string {
	if remainingLives == 0 {
		return hangmanDraw[lives]
	}
	return hangmanDraw[lives-remainingLives+1]
}
