package hangman

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func RandomWord(filename string) string {
	words := FileToStringArray(filename)
	randomIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(words))
	return words[randomIndex]
}

func FileToStringArray(filename string) []string {
	var response []string
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Error when opening file")
	}
	response = strings.Split(string(content), "\n")
	return response
}
