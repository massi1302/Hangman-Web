package main

import (
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

// Variables globales
var templates *template.Template

// validateurs
func init() {
	var err error
	templates, err = template.ParseGlob(TEMPLATE_DIR)
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v", err)
	}
}

// Gestionnaires HTTP
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Validate the form

// Fonctions utilitaires

func setupRoutes() {
	// Serveur de fichiers statiques
	fs := http.FileServer(http.Dir(ASSETS_DIR))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Routes

	http.HandleFunc("/home", homeHandler)
}

func main() {
	setupRoutes()
	log.Printf("Serveur démarré sur http://localhost%s", PORT)
	log.Fatal(http.ListenAndServe("localhost"+PORT, nil))
}
