package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var (
	hangmanDir = filepath.Join(os.TempDir(), "hangman")
	usersDir   = filepath.Join(hangmanDir, "users")
)

func init() {
	if err := os.MkdirAll(hangmanDir, os.ModePerm); err != nil {
		fmt.Println("Erreur lors de la création du dossier du jeu :", err)
		return
	}

	if err := os.MkdirAll(usersDir, os.ModePerm); err != nil {
		fmt.Println("Erreur lors de la création du dossier du jeu :", err)
		return
	}
}

func SaveGameState(username string, gameState *GameState) {
	userFilePath := filepath.Join(usersDir, username+".json")
	file, err := os.Create(userFilePath)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erreur lors de la fermeture du fichier :", err)
		}
	}(file)

	if err = json.NewEncoder(file).Encode(gameState); err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}

func LoadGameState(username string) (*GameState, error) {
	userFilePath := filepath.Join(usersDir, username+".json")
	file, err := os.Open(userFilePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erreur lors de la fermeture du fichier :", err)
		}
	}(file)

	var userGameState *GameState
	if err := json.NewDecoder(file).Decode(&userGameState); err != nil {
		return nil, err
	}

	return userGameState, nil
}
