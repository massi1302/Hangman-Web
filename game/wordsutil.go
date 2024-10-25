package hangman

import (
	"fmt"
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
	var res []string
	var content []byte
	var err error
	if FileExists("Data\\" + filename) {
		content, err = os.ReadFile("Data\\" + filename)
	} else {
		content, err = os.ReadFile(filename)
	}
	if err != nil {
		fmt.Println("Error when opening file", err)
		res = strings.Split(string(content), "\n")
	}
	return res
}
