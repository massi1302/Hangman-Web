package main

import (
	"fmt"
	hangman "hangman/game"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT         = ":8080"
	TEMPLATE_DIR = "./Templates/*.html"
	ASSETS_DIR   = "assets"
	DATE_FORMAT  = "2006-01-02"
)

// Structures de données
type Session struct {
	Username string
}

// Variables globales
var templates *template.Template
var selectedDifficulty string

// validateurs
func init() {
	var err error
	templates, err = template.ParseGlob(TEMPLATE_DIR)
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v", err)
	}
}

// Gestionnaires HTTP
func startGameHandler(respWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Redirect(respWriter, req, "/index", http.StatusSeeOther)
		return
	}

	username := req.FormValue("username")
	if username == "" {
		http.Redirect(respWriter, req, "/index", http.StatusSeeOther)
		return
	}

	// Set username cookie
	cookie := &http.Cookie{
		Name:     "username",
		Value:    username,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	http.SetCookie(respWriter, cookie)

	// Redirect to home page
	http.Redirect(respWriter, req, "/home", http.StatusSeeOther)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func gameHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fmt.Println("Query :", req.URL.Query())
		difficulty := req.FormValue("difficulty")
		usernameCookie, _ := req.Cookie("username")
		filteredGameState := hangman.NewGame(usernameCookie.Value, difficulty)
		if err := templates.ExecuteTemplate(resp, "game", filteredGameState); err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func endgameHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "endgame", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func scoresHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "scores", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setDifficultyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		selectedDifficulty = r.FormValue("difficulty")
		fmt.Println("Difficulté choisie:", selectedDifficulty)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func letterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		usernameCookie, _ := r.Cookie("username")
		letter := r.URL.Query().Get("letter")
		newGameState := hangman.GetGameState(usernameCookie.Value).GuessLetter(letter)
		if !newGameState.GameOver {
			if err := templates.ExecuteTemplate(w, "game", newGameState.ToFilteredGameState()); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if newGameState.Victory {

		} else {
			http.Redirect(w, r, "/endgame", http.StatusSeeOther)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Validate the form

// Fonctions utilitaires

func serveMux() *http.ServeMux {
	fs := http.FileServer(http.Dir(ASSETS_DIR))
	serveMux := http.NewServeMux()

	// Serveur de fichiers statiques
	serveMux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Routes
	serveMux.HandleFunc("/", indexHandler)
	serveMux.HandleFunc("/home", homeHandler)
	serveMux.HandleFunc("/game", gameHandler)
	serveMux.HandleFunc("/endgame", endgameHandler)
	serveMux.HandleFunc("/scores", scoresHandler)
	serveMux.HandleFunc("/index", indexHandler)
	serveMux.HandleFunc("/start-game", startGameHandler)
	serveMux.HandleFunc("/set-difficulty", setDifficultyHandler)
	serveMux.HandleFunc("/guess", letterHandler)

	return serveMux
}

func main() {
	log.Printf("Serveur démarré sur http://localhost%s", PORT)

	log.Fatal(http.ListenAndServe("localhost"+PORT, serveMux()))
}
