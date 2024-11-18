package server

import (
	"fmt"
	"hangman/config"
	"hangman/data"
	"hangman/game"
	"hangman/game/state"
	"html/template"
	"net/http"
)

var templates *template.Template

func addOne(i int) int {
	return i + 1
}

func init() {
	templates = template.Must(template.New("hangman").Funcs(template.FuncMap{"addOne": addOne}).ParseGlob(config.App.Server.StaticWeb.Template.Dir))
}

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := templates.ExecuteTemplate(responseWriter, "index", nil); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	userCookie, err := request.Cookie("username")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}

	gameState, err := state.LoadGameState(userCookie.Value)
	if err != nil {
		fmt.Println("Erreur lors du chargement des gameStates")
		if err := templates.ExecuteTemplate(responseWriter, "home", nil); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := templates.ExecuteTemplate(responseWriter, "home", gameState.ToFilteredGameState()); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	}
}

func scoresHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := templates.ExecuteTemplate(responseWriter, "scores", game.GetScoresManager()); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func gameHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}

	difficulty := request.FormValue("difficulty")
	userCookie, err := request.Cookie("username")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	gameState, err := game.NewGame(userCookie.Value, difficulty)
	state.SaveGameState(userCookie.Value, gameState)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(responseWriter, "game", data.NewGameData(gameState)); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func guessHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
	userCookie, err := request.Cookie("username")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}

	letter := request.URL.Query().Get("letter")
	newGameState := game.GetGameState(userCookie.Value).GuessLetter(letter)
	defer state.SaveGameState(userCookie.Value, newGameState)
	if !newGameState.GameOver {
		if err := templates.ExecuteTemplate(responseWriter, "game", data.NewGameData(newGameState)); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if !newGameState.Victory {
			game.AddScore(userCookie.Value, newGameState.Score, newGameState.Victory)
		}
		if err := templates.ExecuteTemplate(responseWriter, "game", data.NewGameOverData(newGameState)); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		newGameState.ResetScore()
	}
}

func loginHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
		http.Redirect(responseWriter, request, "/index", http.StatusSeeOther)
		return
	}

	username := request.FormValue("username")
	if username == "" {
		http.Redirect(responseWriter, request, "/index", http.StatusSeeOther)
		return
	}

	http.SetCookie(responseWriter, &http.Cookie{Name: "username", Value: username, Path: "/", MaxAge: 3600, HttpOnly: true})
	http.Redirect(responseWriter, request, "/home", http.StatusSeeOther)
}

func continueHandler(responseWriter http.ResponseWriter, request *http.Request) {
	userCookie, err := request.Cookie("username")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	gameState := game.Continue(userCookie.Value)
	if err := templates.ExecuteTemplate(responseWriter, "game", data.NewGameData(gameState)); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func ServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(config.App.Server.StaticWeb.Assets.Dir))
	serveMux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	serveMux.HandleFunc("/", indexHandler)
	serveMux.HandleFunc("/index", indexHandler)
	serveMux.HandleFunc("/login", loginHandler)
	serveMux.HandleFunc("/home", homeHandler)
	serveMux.HandleFunc("/scores", scoresHandler)
	serveMux.HandleFunc("/game", gameHandler)
	serveMux.HandleFunc("/guess", guessHandler)
	serveMux.HandleFunc("/continue", continueHandler)

	return serveMux
}
