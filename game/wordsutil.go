package game

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func RandomWord(path string) (*string, error) {
	words, err := FileToStringArray(path)
	if err != nil {
		return nil, err
	}
	randomIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(words))
	return &words[randomIndex], nil
}

func FileToStringArray(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return strings.Split(string(content), "\r\n"), nil
}
