package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Score struct {
	Username string    `json:"username"`
	Points   int       `json:"points"`
	Victory  bool      `json:"victory"`
	Date     time.Time `json:"date"`
}

type ScoreManager struct {
	HighScores []Score `json:"highScores"`
	History    []Score `json:"history"`
}

const (
	maxHighScores    = 3
	maxHistoryScores = 5
)

var (
	scoresFile = filepath.Join(hangmanDir, "scores.json")
)

func NewScoreManager() *ScoreManager {
	return &ScoreManager{
		HighScores: make([]Score, 0),
		History:    make([]Score, 0),
	}
}

func (sm *ScoreManager) LoadScores() error {
	data, err := os.ReadFile(scoresFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Si le fichier n'existe pas, on crée un nouveau score manager
			return sm.SaveScores()
		}
		return fmt.Errorf("error reading scores file: %v", err)
	}

	err = json.Unmarshal(data, sm)
	if err != nil {
		return fmt.Errorf("error unmarshaling scores: %v", err)
	}

	return nil
}

func (sm *ScoreManager) SaveScores() error {
	data, err := json.MarshalIndent(sm, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling scores: %v", err)
	}

	err = os.WriteFile(scoresFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing scores file: %v", err)
	}

	return nil
}

func (sm *ScoreManager) AddScore(username string, points int, victory bool) error {
	newScore := Score{
		Username: username,
		Points:   points,
		Victory:  victory,
		Date:     time.Now(),
	}

	// Ajouter à l'historique
	sm.History = append(sm.History, newScore)
	if len(sm.History) > maxHistoryScores {
		sm.History = sm.History[len(sm.History)-maxHistoryScores:]
	}

	// Vérifier si le score mérite d'être dans les high scores
	if points > 0 {
		sm.HighScores = append(sm.HighScores, newScore)
	}

	// Trier les high scores par points décroissants
	sort.Slice(sm.HighScores, func(i, j int) bool {
		return sm.HighScores[i].Points > sm.HighScores[j].Points
	})

	// Garder uniquement les meilleurs scores
	if len(sm.HighScores) > maxHighScores {
		sm.HighScores = sm.HighScores[:maxHighScores]
	}

	return sm.SaveScores()
}

// Fonction utilitaire pour formater la date pour l'affichage
func (s Score) FormattedDate() string {
	return s.Date.Format("02/01/2006 15:04")
}

// Fonction utilitaire pour formater le résultat pour l'affichage
func (s Score) FormattedResult() string {
	if s.Victory {
		return "Victoire"
	}
	return "Défaite"
}
