package state

import (
	"html/template"
	"slices"
	"strings"
)

// Lives représente le nombre de vies
const Lives = 10

var hangmanDraw = make(map[int]string)

// LoadGameStates loads all game states for a given user.

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

// GameState représente l'état actuel d'une partie
type GameState struct {
	Difficulty    string   `json:"difficulty"`
	Word          string   `json:"word"`
	DisplayedWord []string `json:"displayedWord"`
	Lives         int      `json:"Lives"`
	UsedLetters   []string `json:"usedLetters"`
	GameOver      bool     `json:"gameOver"`
	Victory       bool     `json:"victory"`
	Score         int      `json:"score"`
	BestScore     int      `json:"bestScore"`
}

// FilteredGameState représente l'état actuel d'une partie filtrée
type FilteredGameState struct {
	Difficulty    string
	DisplayedWord []string
	Lives         int
	UsedLetters   []string
	GameOver      bool
	Victory       bool
	Score         int
	BestScore     int
}

// ToFilteredGameState convertit un GameState en FilteredGameState
func (g *GameState) ToFilteredGameState() *FilteredGameState {
	return &FilteredGameState{
		Difficulty:    g.Difficulty,
		DisplayedWord: g.DisplayedWord,
		Lives:         g.Lives,
		UsedLetters:   g.UsedLetters,
		GameOver:      g.GameOver,
		Victory:       g.Victory,
		Score:         g.Score,
		BestScore:     g.BestScore,
	}
}

// GuessLetter traite une tentative de lettre
func (g *GameState) GuessLetter(letter string) *GameState {
	letter = strings.ToLower(letter)

	for _, usedLetter := range g.UsedLetters {
		if usedLetter == letter {
			return g
		}
	}

	g.UsedLetters = append(g.UsedLetters, letter)

	correct := false
	for i, char := range g.Word {
		if string(char) == letter {
			g.DisplayedWord[i] = letter
			correct = true
		}
	}

	if !correct {
		g.Lives--
	}

	g.isEndGame()

	return g
}

// GuessWord traite une tentative de mot
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
	g.isEndGame()
	return false
}

// isEndGame vérifie si la partie est terminée
func (g *GameState) isEndGame() bool {
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

func (g *GameState) ResetScore() {
	if g.Lives <= 0 {
		if g.BestScore < g.Score {
			g.BestScore = g.Score
		}
		g.Score = 0
	}
}

// IsUsedLetter vérifie si la lettre a déjà été utilisée
func (f *FilteredGameState) IsUsedLetter(letter string) bool {
	return slices.Contains(f.UsedLetters, strings.ToLower(letter)) && !f.IsWordsLetter(letter)
}

func (f *FilteredGameState) IsWordsLetter(letter string) bool {
	return slices.Contains(f.DisplayedWord, strings.ToLower(letter))
}

// DrawHearts dessine les coeurs
func (f *FilteredGameState) DrawHearts() template.HTML {
	heartIcon := "❤️"
	var heartString string
	for i := 0; i < f.Lives; i++ {
		heartString += heartIcon + " "
	}
	return template.HTML(heartString)
}

// DrawHangman dessine le pendu
func (f *FilteredGameState) DrawHangman(remainingLives int) string {
	if remainingLives == 0 {
		return hangmanDraw[Lives]
	}
	return hangmanDraw[Lives-remainingLives+1]
}
