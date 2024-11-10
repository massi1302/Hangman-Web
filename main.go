package main

import (
	"fmt"
	"hangman/config"
	hangman "hangman/game"
	"hangman/util"
	"html/template"
	"log"
	"net"
	"net/http"
)

// Variables globales
var templates *template.Template
var selectedDifficulty string

// validateurs
func init() {
	var err error
	templates, err = template.ParseGlob(config.App.Server.StaticWeb.Template.Dir)
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v\n", err)
	}
}

// Gestionnaires HTTP
func loginHandler(respWriter http.ResponseWriter, req *http.Request) {
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
		difficulty := req.FormValue("difficulty")
		usernameCookie, err := req.Cookie("username")
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		filteredGameState, err := hangman.NewGame(usernameCookie.Value, difficulty)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templates.ExecuteTemplate(resp, "game", util.NewGameData(filteredGameState)); err != nil {
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

		if err := templates.ExecuteTemplate(w, "game", util.NewGameData(newGameState.ToFilteredGameState())); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Validate the form

// Fonctions utilitaires

func serveMux() *http.ServeMux {
	fs := http.FileServer(http.Dir(config.App.Server.StaticWeb.Assets.Dir))
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
	serveMux.HandleFunc("/login", loginHandler)
	serveMux.HandleFunc("/set-difficulty", setDifficultyHandler)
	serveMux.HandleFunc("/guess", letterHandler)

	return serveMux
}

func main() {
	log.Printf("Serveur démarré sur http://%s\n", net.JoinHostPort(config.App.Server.URL, config.App.Server.Port))
	log.Fatal(http.ListenAndServe(net.JoinHostPort(config.App.Server.URL, config.App.Server.Port), serveMux()))
}
